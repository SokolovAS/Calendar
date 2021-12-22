package calendar

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	jwtWrapper := JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	serv := NewAuthService()

	_, err := serv.GenerateToken("email@test.com", &jwtWrapper)
	if err != nil {
		t.Errorf("error %s", err)
	}
}

func TestValidate(t *testing.T) {
	jwtWrapper := JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	serv := NewAuthService()
	newToken, err := serv.GenerateToken("email@test.com", &jwtWrapper)

	_, err = serv.Validate("Bearer " + newToken)
	if err != nil {
		t.Errorf("error %s", err)
	}

	testToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InBvc3RtYW5AZ21haWwuY29tIiwiZXhwIjoxNjM5NzY2MzY2LCJpc3MiOiJBdXRoU2VydmljZSJ9.wA0eqkUacIN0dxByR3A9JsXZsTVDbTmGndMqKx8_3Sc"
	_, err = serv.Validate("Bearer " + testToken)
	require.Error(t, err)
}
