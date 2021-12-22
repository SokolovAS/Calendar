package http

import (
	"Calendar/entity"
	"Calendar/internal/services/calendar"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewAuthHandlerMock() *AuthHandler {
	return &AuthHandler{
		authS: &authServiceMock{},
		userS: &userServiceMock{},
	}
}

const testEmail = "email@gmail.com"

type userServiceMock struct{}

func (*userServiceMock) CreateUserRecord(m entity.User) error {
	return nil
}
func (*userServiceMock) HashPassword(m *entity.User) error {
	return nil
}
func (*userServiceMock) CheckPassword(providedPassword string, userPassword string) error {
	return nil
}
func (*userServiceMock) GetEmail(email string) (entity.User, error) {
	return entity.User{}, nil
}

type authServiceMock struct{}

func (m *authServiceMock) GenerateToken(email string, j *calendar.JwtWrapper) (signedToken string, err error) {
	return
}

func (m *authServiceMock) ValidateToken(signedToken string, j *calendar.JwtWrapper) (claims *calendar.JwtClaim, err error) {
	return
}

func (*authServiceMock) Validate(clientToken string) (string, error) {
	_ = clientToken
	return testEmail, nil
}

func NewAuthServiceMock() AuthService {
	return &authServiceMock{}
}

func TestSignup(t *testing.T) {
	var jsonStr = []byte(`{"email":"somail@gmail.com"}`)
	req, err := http.NewRequest("POST", "/http://localhost:8000/signup", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	h := NewAuthHandlerMock()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Signup)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestLogin(t *testing.T) {
	var jsonStr = []byte(`{"email":"somail@gmail.com"}`)
	req, err := http.NewRequest("POST", "/http://localhost:8000/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	h := NewAuthHandlerMock()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Login)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
