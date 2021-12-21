package http

import (
	"Calendar/entity"
	"Calendar/internal/services/calendar"
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAll(t *testing.T) {

	testCases := []struct {
		name    string
		mock    func(ctrl *gomock.Controller) calendar.EventService
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
			mock: func(ctrl *gomock.Controller) calendar.EventService {
				events := []entity.Event{
					{"1", "Title1", "Description1", "DateTiem", "Duration1", "Notes1"},
				}
				eS := NewMockEventService(ctrl)

				eS.EXPECT().GetAll().Return(events, nil)
				return eS
			},
		},
		{
			name:    "bad path",
			wantErr: true,
			mock: func(ctrl *gomock.Controller) calendar.EventService {

				eS := NewMockEventService(ctrl)

				eS.EXPECT().GetAll().Return([]entity.Event{}, errors.New("error getting events"))
				return eS
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()

			eS := tc.mock(ctl)
			eH := eventHandler{
				eServ: eS,
			}

			req, _ := http.NewRequest("GET", "http://localhost:8000/events", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InBvc3RtYW5AZ21haWwuY29tIiwiZXhwIjoxNjM5NzY2MzY2LCJpc3MiOiJBdXRoU2VydmljZSJ9.wA0eqkUacIN0dxByR3A9JsXZsTVDbTmGndMqKx8_3Sc")

			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(eH.GetAll)

			handler.ServeHTTP(rr, req)

			if tc.wantErr {
				if status := rr.Code; status != http.StatusInternalServerError {
					t.Errorf("expected an error: got %v want %v",
						status, http.StatusInternalServerError)
				}
				return
			}
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

		})
	}

	//e := entity.Event{}
	//hz:= io.Reader(rr.Body)
	//fmt.Println("SOME DAATA", hz)
	//err := json.NewDecoder(hz).Decode(&e)
	//if err != nil{
	//	t.Errorf("Unexpected error while unmarshaling! %s", err)
	//}
	//
	//if events[0].Id != e.Id {
	//	t.Errorf("got %s want %s", e.Id, events[0].Id)
	//}

}

func TestGetOne(t *testing.T) {
	testCases := []struct {
		name    string
		mock    func(ctrl *gomock.Controller) calendar.EventService
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
			mock: func(ctrl *gomock.Controller) calendar.EventService {

				eS := NewMockEventService(ctrl)
				eS.EXPECT().GetOne(gomock.Any()).Return(entity.Event{Id: "1"}, nil)

				return eS
			},
		},
		{
			name:    "bad path",
			wantErr: true,
			mock: func(ctrl *gomock.Controller) calendar.EventService {

				eS := NewMockEventService(ctrl)
				eS.EXPECT().GetOne(gomock.Any()).Return(entity.Event{}, errors.New("error getting event"))

				return eS
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()

			eS := tc.mock(ctl)
			eH := eventHandler{
				eServ: eS,
			}

			var jsonStr = []byte(`{"id": "1"}`)
			req, _ := http.NewRequest("GET", "http://localhost:8000/event", bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InBvc3RtYW5AZ21haWwuY29tIiwiZXhwIjoxNjM5NzY2MzY2LCJpc3MiOiJBdXRoU2VydmljZSJ9.wA0eqkUacIN0dxByR3A9JsXZsTVDbTmGndMqKx8_3Sc")

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(eH.GetOne)
			handler.ServeHTTP(rr, req)

			if tc.wantErr {
				if status := rr.Code; status != http.StatusInternalServerError {
					t.Errorf("expected an error: got %v want %v",
						status, http.StatusInternalServerError)
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

func TestUpdate(t *testing.T) {
	testCases := []struct {
		name    string
		mock    func(ctrl *gomock.Controller) calendar.EventService
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
			mock: func(ctrl *gomock.Controller) calendar.EventService {

				eS := NewMockEventService(ctrl)
				eS.EXPECT().Update(gomock.Any()).Return(entity.Event{Id: "1", Title: "Title updated"}, nil)

				return eS
			},
		},
		{
			name:    "bad path",
			wantErr: true,
			mock: func(ctrl *gomock.Controller) calendar.EventService {

				eS := NewMockEventService(ctrl)
				eS.EXPECT().Update(gomock.Any()).Return(entity.Event{}, errors.New("error updating event"))

				return eS
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()

			eS := tc.mock(ctl)
			eH := eventHandler{
				eServ: eS,
			}

			var jsonStr = []byte(`{"id": "1", "title":"Title updated"}`)
			req, _ := http.NewRequest("PUT", "http://localhost:8000/event", bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InBvc3RtYW5AZ21haWwuY29tIiwiZXhwIjoxNjM5NzY2MzY2LCJpc3MiOiJBdXRoU2VydmljZSJ9.wA0eqkUacIN0dxByR3A9JsXZsTVDbTmGndMqKx8_3Sc")

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(eH.Update)
			handler.ServeHTTP(rr, req)

			if tc.wantErr {
				if status := rr.Code; status != http.StatusInternalServerError {
					t.Errorf("expected an error: got %v want %v",
						status, http.StatusInternalServerError)
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

func TestAdd(t *testing.T) {
	testCases := []struct {
		name    string
		mock    func(ctrl *gomock.Controller) calendar.EventService
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
			mock: func(ctrl *gomock.Controller) calendar.EventService {

				eS := NewMockEventService(ctrl)
				eS.EXPECT().Add(entity.Event{Id: "2", Title: "Title NEW"}).Return(entity.Event{Id: "2", Title: "Title NEW"}, nil)

				return eS
			},
		},
		{
			name:    "bad path",
			wantErr: true,
			mock: func(ctrl *gomock.Controller) calendar.EventService {

				eS := NewMockEventService(ctrl)
				eS.EXPECT().Add(entity.Event{Id: "2", Title: "Title NEW"}).Return(entity.Event{}, errors.New("error adding event"))

				return eS
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()

			eS := tc.mock(ctl)
			eH := eventHandler{
				eServ: eS,
			}

			var jsonStr = []byte(`{"id": "2", "title":"Title NEW"}`)
			req, _ := http.NewRequest("POST", "http://localhost:8000/event", bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InBvc3RtYW5AZ21haWwuY29tIiwiZXhwIjoxNjM5NzY2MzY2LCJpc3MiOiJBdXRoU2VydmljZSJ9.wA0eqkUacIN0dxByR3A9JsXZsTVDbTmGndMqKx8_3Sc")

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(eH.Add)
			handler.ServeHTTP(rr, req)

			if tc.wantErr {
				if status := rr.Code; status != http.StatusInternalServerError {
					t.Errorf("expected an error: got %v want %v",
						status, http.StatusInternalServerError)
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

func TestDelete(t *testing.T) {
	testCases := []struct {
		name    string
		mock    func(ctrl *gomock.Controller) calendar.EventService
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
			mock: func(ctrl *gomock.Controller) calendar.EventService {

				eS := NewMockEventService(ctrl)
				eS.EXPECT().Delete(gomock.Any())

				return eS
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()

			eS := tc.mock(ctl)
			eH := eventHandler{
				eServ: eS,
			}

			var jsonStr = []byte(`{"id": "2"`)
			req, _ := http.NewRequest("DELETE", "http://localhost:8000/event", bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InBvc3RtYW5AZ21haWwuY29tIiwiZXhwIjoxNjM5NzY2MzY2LCJpc3MiOiJBdXRoU2VydmljZSJ9.wA0eqkUacIN0dxByR3A9JsXZsTVDbTmGndMqKx8_3Sc")

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(eH.Delete)
			handler.ServeHTTP(rr, req)

			if tc.wantErr {
				if status := rr.Code; status != http.StatusInternalServerError {
					t.Errorf("expected an error: got %v want %v",
						status, http.StatusInternalServerError)
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
