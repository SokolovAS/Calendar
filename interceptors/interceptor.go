package interceptors

import (
	"Calendar/internal/server/http"
	"Calendar/internal/services/calendar"
	"context"
	_ "github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

type AuthMD struct {
	As http.AuthService
}

const (
	authHeader = "authorization"
	bearerAuth = "bearer"
)

func (a *AuthMD) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		credentials := a.getAuthCredentials(ctx)
		if credentials == "" {
			return nil, status.Error(codes.Unauthenticated, "unauthenticated")
		}

		jwtWrapper := calendar.JwtWrapper{
			SecretKey: "verysecretkey",
			Issuer:    "AuthService",
		}

		claim, err := a.As.ValidateToken(credentials, &jwtWrapper)

		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "incorrect auth token")
		}

		ctx = context.WithValue(ctx, "user", claim.Email)
		return handler(ctx, req)
	}
}

func (a *AuthMD) getAuthCredentials(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	values := md.Get(authHeader)
	if len(values) == 0 {
		return ""
	}

	fields := strings.SplitN(values[0], " ", 2)
	if len(fields) < 2 {
		return ""
	}

	if !strings.EqualFold(fields[0], bearerAuth) {
		return ""
	}
	return fields[1]
}
