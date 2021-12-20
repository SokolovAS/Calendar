package http

import (
	"Calendar/entity"
	"bytes"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAll(t *testing.T) {
	events := []entity.Event{
		{"1", "Title1", "Description1", "DateTiem", "Duration1", "Notes1"},
	}
	//eJson, _ := json.Marshal(events)

	ctrl := gomock.NewController(t)
	eS := NewMockEventService(ctrl)

	eS.EXPECT().GetAll().Return(events, nil)

	eH := eventHandler{
		eServ: eS,
	}

	req, _ := http.NewRequest("GET", "http://localhost:8000/events", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InBvc3RtYW5AZ21haWwuY29tIiwiZXhwIjoxNjM5NzY2MzY2LCJpc3MiOiJBdXRoU2VydmljZSJ9.wA0eqkUacIN0dxByR3A9JsXZsTVDbTmGndMqKx8_3Sc")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(eH.GetAll)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
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
	ctrl := gomock.NewController(t)
	eS := NewMockEventService(ctrl)

	eS.EXPECT().GetOne(gomock.Any()).Return(entity.Event{Id: "1"}, nil)

	eH := eventHandler{
		eServ: eS,
	}

	var jsonStr = []byte(`{"id": "2","title": "Title2 updated","description": "Description2 updated","dateTime": "DateTiem2 updated","duration": "Duration2 updated","notes": "Notes2 updated"}`)
	req, _ := http.NewRequest("GET", "http://localhost:8000/events", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InBvc3RtYW5AZ21haWwuY29tIiwiZXhwIjoxNjM5NzY2MzY2LCJpc3MiOiJBdXRoU2VydmljZSJ9.wA0eqkUacIN0dxByR3A9JsXZsTVDbTmGndMqKx8_3Sc")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(eH.GetOne)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
