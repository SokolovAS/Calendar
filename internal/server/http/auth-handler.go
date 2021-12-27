package http

import (
	"Calendar/entity"
	"Calendar/internal/services/calendar"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type AuthService interface {
	GenerateToken(id string, email string, j *calendar.JwtWrapper) (signedToken string, err error)
	ValidateToken(signedToken string, j *calendar.JwtWrapper) (claims *calendar.JwtClaim, err error)
	Validate(clientToken string) (*calendar.JwtClaim, error)
}

type EventService interface {
	GetAll() ([]entity.Event, error)
	GetOne(id string) (entity.Event, error)
	Add(event entity.Event) (entity.Event, error)
	Update(event entity.Event) (entity.Event, error)
	Delete(id string) error
}

type UserService interface {
	CreateUserRecord(m entity.User) error
	HashPassword(m *entity.User) error
	CheckPassword(providedPassword string, userPassword string) error
	GetEmail(email string) (entity.User, error)
}

type AuthHandler struct {
	authS AuthService
	userS UserService
}

func NewAuthHandler(aS AuthService, uS UserService) *AuthHandler {
	return &AuthHandler{
		authS: aS,
		userS: uS,
	}
}

// Signup creates a user in db
func (aH *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		assertMarshalingError(w, err)
	}

	err = aH.userS.HashPassword(&user)
	if err != nil {
		log.Println(err.Error())
		assertMarshalingError(w, err)
	}

	err = aH.userS.CreateUserRecord(user)
	if err != nil {
		log.Println(err)
		assertMarshalingError(w, err)
	}

	result, err := json.Marshal(user)
	assertMarshalingError(w, err)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(result)
	assertResponseError(err)
}

// LoginPayload login body
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse token response
type LoginResponse struct {
	Token string `json:"token"`
}

// Login logs users in
func (aH *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	fmt.Println("BODY", err)
	if err != nil {
		log.Println(err)
		assertMarshalingError(w, err)
	}

	user, err := aH.userS.GetEmail(payload.Email)

	if err != nil {
		assertError(w, `"error":"Error fetching data"`)
	}

	err = aH.userS.CheckPassword(payload.Password, user.Password)
	if err != nil {
		var tge *calendar.ServiceErr
		if errors.As(err, &tge) {
			w.WriteHeader(tge.Code)
			mess := fmt.Sprintf("%#v\n", tge)
			write, err := w.Write([]byte(mess))
			if err != nil {
				return
			}
			_ = write
		}
		assertError(w, `"error":"Error password"`)
	}

	jwtWrapper := calendar.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := aH.authS.GenerateToken(strconv.Itoa(int(user.ID)), user.Email, &jwtWrapper)
	if err != nil {
		var tge *calendar.ServiceErr
		if errors.As(err, &tge) {
			w.WriteHeader(tge.Code)
			mess := fmt.Sprintf("%#v\n", tge)
			write, err := w.Write([]byte(mess))
			if err != nil {
				return
			}
			_ = write
		}
		assertError(w, `"msg": "error signing token"`)
	}

	tokenResponse := LoginResponse{
		Token: signedToken,
	}

	res, err := json.Marshal(tokenResponse)
	assertMarshalingError(w, err)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	assertResponseError(err)
}
