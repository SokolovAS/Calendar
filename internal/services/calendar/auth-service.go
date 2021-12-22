package calendar

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type AuthService struct {
	Claim JwtClaim
}

func NewAuthService(r SqliteRepo) *AuthService {
	return &AuthService{}
}

// JwtWrapper wraps the signing key and the issuer
type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// JwtClaim adds email as a claim to the token
type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

// GenerateToken generates a jwt token
func (*AuthService) GenerateToken(email string, j *JwtWrapper) (signedToken string, err error) {
	claims := &JwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}

	return
}

//ValidateToken validates the jwt token
func (*AuthService) ValidateToken(signedToken string, j *JwtWrapper) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return
	}
	return
}

func (aS *AuthService) Validate(clientToken string) (string, error) {
	if clientToken == "" {
		return "", errors.New(`"error":"No Authorization header provided"`)
	}

	extractedToken := strings.Split(clientToken, "Bearer ")

	if len(extractedToken) == 2 {
		clientToken = strings.TrimSpace(extractedToken[1])
	} else {
		return "", errors.New(`"error":"Incorrect Format of Authorization Token`)
	}

	jwtWrapper := JwtWrapper{
		SecretKey: "verysecretkey",
		Issuer:    "AuthService",
	}

	claims, err := aS.ValidateToken(clientToken, &jwtWrapper)
	if err != nil {
		return "", errors.New(`"error":"Error token validation"`)
	}

	return claims.Email, err
}
