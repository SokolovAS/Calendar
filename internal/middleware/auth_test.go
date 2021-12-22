package middleware

import (
	http2 "Calendar/internal/server/http"
	"Calendar/internal/services/calendar"
	"github.com/gorilla/mux"
	"net/http"
	"testing"
)

const testEmail = "test@gmail.com"

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

func NewAuthServiceMock() http2.AuthService {
	return &authServiceMock{}
}

func NewMiddlewareMock(aS AuthService) *Middleware {
	return &Middleware{
		authS: aS,
	}
}

func TestAuthz(t *testing.T) {
	fn1 := func(w http.ResponseWriter, r *http.Request) {
		clientToken := r.Header.Get("Authorization")
		if clientToken == "" {
			t.Errorf("No Authorization header provided")
		}
		params := mux.Vars(r)
		email := params["email"]
		if email != testEmail {
			t.Errorf("expected %s get %s", testEmail, email)
		}
	}

	aSmock := NewAuthServiceMock()
	mid := NewMiddlewareMock(aSmock)
	h1 := mid.Authz(http.HandlerFunc(fn1))
	req, _ := http.NewRequest("GET", "http://localhost:8000/events", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InBvc3RtYW5AZ21haWwuY29tIiwiZXhwIjoxNjM5NzY2MzY2LCJpc3MiOiJBdXRoU2VydmljZSJ9.wA0eqkUacIN0dxByR3A9JsXZsTVDbTmGndMqKx8_3Sc")
	h1.ServeHTTP(nil, req)
}
