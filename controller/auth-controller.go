package controller

import (
	"Calendar/entity"
	database "Calendar/initdb.d"
	auth "Calendar/internal/middleware"
	"Calendar/internal/services/calendar"
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var (
	UserService calendar.UserService = calendar.NewUserService()
)

type authController struct{}

type AuthController interface {
	Signup(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

func NewAuthController() AuthController {
	return &authController{}
}

// Signup creates a user in db
func (*authController) Signup(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		assertMarshalingError(w, err)
	}

	err = UserService.HashPassword(&user)
	if err != nil {
		log.Println(err.Error())
		assertMarshalingError(w, err)
	}

	err = UserService.CreateUserRecord(user)
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
func (*authController) Login(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayload
	var user entity.User

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		assertMarshalingError(w, err)
	}

	result := database.GlobalDB.Where("email = ?", payload.Email).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		assertGormError(w, `"error":"Error fetching data"`)
	}

	err = UserService.CheckPassword(payload.Password, user.Password)
	if err != nil {
		log.Println(err)
		assertGormError(w, `"error":"Error password"`)
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)
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
