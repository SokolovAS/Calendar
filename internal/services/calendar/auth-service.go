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

func NewAuthService() *AuthService {
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
	Id    string
	Email string
	jwt.StandardClaims
}

// GenerateToken generates a jwt token
func (*AuthService) GenerateToken(id string, email string, j *JwtWrapper) (signedToken string, err error) {
	claims := &JwtClaim{
		Id:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		e := &ServiceErr{500, "Error token generation"}
		return "", e
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
		e := &ServiceErr{401, "Error token validation"}
		return claims, e
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		e := &ServiceErr{403, "Couldn't parse claims"}
		return claims, e
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		e := &ServiceErr{401, "JWT is expired"}
		return claims, e
	}
	return
}

func (aS *AuthService) Validate(clientToken string) (*JwtClaim, error) {
	if clientToken == "" {
		return &JwtClaim{}, errors.New(`"error":"No Authorization header provided"`)
	}

	extractedToken := strings.Split(clientToken, "Bearer ")

	if len(extractedToken) == 2 {
		clientToken = strings.TrimSpace(extractedToken[1])
	} else {
		e := &ServiceErr{401, "Incorrect Format of Authorization Token"}
		return &JwtClaim{}, e
	}

	jwtWrapper := JwtWrapper{
		SecretKey: "verysecretkey",
		Issuer:    "AuthService",
	}

	claims, err := aS.ValidateToken(clientToken, &jwtWrapper)
	if err != nil {
		e := &ServiceErr{401, "Incorrect Format of Authorization Token"}
		return &JwtClaim{}, e
	}

	return claims, err
}
