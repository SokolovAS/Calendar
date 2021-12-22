package calendar

import (
	"Calendar/internal/repository"
	"testing"
)

func TestGenerateToken(t *testing.T) {

	jwtWrapper := JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	repo := repository.NewSqliteRepo()
	serv := NewAuthService(repo)

	_, err := serv.GenerateToken("email@test.com", &jwtWrapper)
	if err != nil {
		t.Errorf("error %s", err)
	}
}

func TestValidate(t *testing.T) {
	testToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InBvc3RtYW5AZ21haWwuY29tIiwiZXhwIjoxNjM5NzY2MzY2LCJpc3MiOiJBdXRoU2VydmljZSJ9.wA0eqkUacIN0dxByR3A9JsXZsTVDbTmGndMqKx8_3Sc"
	repo := repository.NewSqliteRepo()
	serv := NewAuthService(repo)
	_, err := serv.Validate(testToken)
	if err != nil {
		t.Errorf("error %s", err)
	}
}
