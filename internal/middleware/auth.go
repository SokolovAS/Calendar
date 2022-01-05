package middleware

import (
	"Calendar/internal/services/calendar"
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type AuthService interface {
	Validate(clientToken string) (*calendar.JwtClaim, error)
}

type Middleware struct {
	authS AuthService
}

func NewMiddleware(a AuthService) *Middleware {
	return &Middleware{
		authS: a,
	}
}

// Authz validates token and authorizes users
func (m *Middleware) Authz(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		id := uuid.New()
		ctx = context.WithValue(ctx, "requestId", id.String())
		r = r.WithContext(ctx)

		clientToken := r.Header.Get("Authorization")
		c, err := m.authS.Validate(clientToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte(`"error":"Incorrect Format of Authorization Token or so`))
			if err != nil {
			}
			return
		}
		r = mux.SetURLVars(r, map[string]string{
			"id":    c.Id,
			"email": c.Email,
		})
		next.ServeHTTP(w, r)
	})
}
