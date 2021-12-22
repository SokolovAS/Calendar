package middleware

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testEmail = "test@gmail.com"

func TestAuthz(t *testing.T) {

	testCases := []struct {
		name    string
		mock    func(ctrl *gomock.Controller) AuthService
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
			mock: func(ctrl *gomock.Controller) AuthService {

				aS := NewMockAuthService(ctrl)
				aS.EXPECT().Validate(gomock.Any()).Return(testEmail, nil)
				return aS
			},
		},
		{
			name:    "bad path",
			wantErr: true,
			mock: func(ctrl *gomock.Controller) AuthService {

				aS := NewMockAuthService(ctrl)
				aS.EXPECT().Validate(gomock.Any()).Return("", errors.New("Error!"))
				return aS
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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

			ctrl := gomock.NewController(t)
			mock := tc.mock(ctrl)

			mid := NewMiddleware(mock)
			h1 := mid.Authz(http.HandlerFunc(fn1))
			req, _ := http.NewRequest("GET", "http://localhost:8000/events", nil)
			req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InBvc3RtYW5AZ21haWwuY29tIiwiZXhwIjoxNjM5NzY2MzY2LCJpc3MiOiJBdXRoU2VydmljZSJ9.wA0eqkUacIN0dxByR3A9JsXZsTVDbTmGndMqKx8_3Sc")
			rr := httptest.NewRecorder()
			h1.ServeHTTP(rr, req)

			if tc.wantErr {
				if status := rr.Code; status == http.StatusOK {
					t.Errorf("handler returned success status code: got %v want error %v",
						status, http.StatusOK)
				}
				return
			}

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
		})

	}

}
