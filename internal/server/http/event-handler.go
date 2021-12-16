package http

import (
	"Calendar/entity"
	"Calendar/internal/services/calendar"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

type eventHandler struct {
	eServ calendar.EventService
	uServ calendar.UserService
}

type EventHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetOne(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewEventHandler() EventHandler {
	return &eventHandler{
		eServ: calendar.NewEventService(),
		uServ: calendar.NewUserService(),
	}
}

func assertMarshalingError(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		write, err := w.Write([]byte(`"error":"Error marshalling data"`))
		if err != nil {
			return
		}
		_ = write
	}
}

func assertGormError(w http.ResponseWriter, error string) {
	w.WriteHeader(http.StatusInternalServerError)
	write, err := w.Write([]byte(error))
	if err != nil {
		return
	}
	_ = write
}

func assertResponseError(err error) {
	if err != nil {
		return
	}
}

func (eventH *eventHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	//---------------------------------------------------------Should be removed
	//var user entity.User
	params := mux.Vars(r)
	email := params["email"]

	_, err := eventH.uServ.GetEmail(email)

	if err == gorm.ErrRecordNotFound {
		assertGormError(w, `"error":"user not found"`)
	}
	//----------------------------------------------------------Should be removed

	events, _ := eventH.eServ.GetAll()

	w.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(events)
	assertMarshalingError(w, err)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(result)
	assertResponseError(err)
}

func (eventH *eventHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var event entity.Event
	err := json.NewDecoder(r.Body).Decode(&event)

	res, err := eventH.eServ.GetOne(event.Id)
	if err != nil {
		fmt.Println(err)
	}

	result, err := json.Marshal(res)
	assertMarshalingError(w, err)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(result)
	assertResponseError(err)
}

func (eventH *eventHandler) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var event entity.Event
	err := json.NewDecoder(r.Body).Decode(&event)

	e, _ := eventH.eServ.Add(event)

	result, err := json.Marshal(e)
	assertMarshalingError(w, err)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(result)
	assertResponseError(err)
}

func (eventH *eventHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var event entity.Event
	err := json.NewDecoder(r.Body).Decode(&event)

	e, _ := eventH.eServ.Update(event)

	result, err := json.Marshal(e)
	assertMarshalingError(w, err)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(result)
	assertResponseError(err)
}

func (eventH *eventHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var event entity.Event
	err := json.NewDecoder(r.Body).Decode(&event)

	eventH.eServ.Delete(event.Id)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(`"result":"ok`))
	assertResponseError(err)
}
