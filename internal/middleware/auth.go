package middleware

import (
	"Calendar/internal/services/calendar"
	"net/http"
)

type Middleware interface {
	Authz(http.Handler) http.Handler
}

type middleware struct {
	authS calendar.AuthService
}

func NewMiddleware() Middleware {
	return &middleware{
		authS: calendar.NewAuthService(),
	}
}

// Authz validates token and authorizes users
func (m *middleware) Authz(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r, err := m.authS.Validate(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte(`"error":"Incorrect Format of Authorization Token or so`))
			if err != nil {
			}
			return
		}
		next.ServeHTTP(w, r)
	})
}
