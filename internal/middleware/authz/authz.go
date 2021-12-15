package middlewares

import (
	"Calendar/internal/middleware"
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

// Authz validates token and authorizes users
func Authz(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		clientToken := r.Header.Get("Authorization")
		if clientToken == "" {
			w.WriteHeader(http.StatusInternalServerError)
			write, err := w.Write([]byte(`"error":"No Authorization header provided"`))
			if err != nil {
			}
			_ = write
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			write, err := w.Write([]byte(`"error":"Incorrect Format of Authorization Token`))
			if err != nil {
			}
			_ = write
			return
		}

		jwtWrapper := auth.JwtWrapper{
			SecretKey: "verysecretkey",
			Issuer:    "AuthService",
		}

		claims, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			write, err := w.Write([]byte(`"error":"Error token validation"`))
			if err != nil {
			}
			_ = write
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "email", claims.Email)

		r = mux.SetURLVars(r, map[string]string{
			"email": claims.Email,
		})

		next.ServeHTTP(w, r)
	})
}
