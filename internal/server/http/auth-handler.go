package http

import (
	"Calendar/entity"
	"Calendar/internal/services/calendar"
	"encoding/json"
	"log"
	"net/http"
)

type authHandler struct {
	authS calendar.AuthService
	userS calendar.UserService
}

type AuthHandler interface {
	Signup(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

func NewAuthHandler() AuthHandler {
	return &authHandler{
		authS: calendar.NewAuthService(),
		userS: calendar.NewUserService(),
	}
}

// Signup creates a user in db
func (aH *authHandler) Signup(w http.ResponseWriter, r *http.Request) {
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
func (aH *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		assertMarshalingError(w, err)
	}

	user, err := aH.userS.GetEmail(payload.Email)

	if err != nil {
		assertGormError(w, `"error":"Error fetching data"`)
	}

	err = aH.userS.CheckPassword(payload.Password, user.Password)
	if err != nil {
		log.Println(err)
		assertGormError(w, `"error":"Error password"`)
	}

	jwtWrapper := calendar.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := aH.authS.GenerateToken(user.Email, &jwtWrapper)
	if err != nil {
		assertGormError(w, `"msg": "error signing token"`)
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
