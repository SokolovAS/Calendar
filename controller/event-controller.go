package controller

import (
	"Calendar/entity"
	database "Calendar/initdb.d"
	"Calendar/internal/services/calendar"
	"Calendar/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

type controller struct{}

type EventController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetOne(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewEventController() EventController {
	return &controller{}
}

var (
	EventService calendar.EventService = calendar.NewEventService()
)

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

func (*controller) GetAll(w http.ResponseWriter, r *http.Request) {
	_ = r

	var user models.User
	params := mux.Vars(r)
	email := params["email"]

	data := database.GlobalDB.Where("email = ?", email).First(&user)

	if data.Error == gorm.ErrRecordNotFound {
		assertGormError(w, `"error":"user not found"`)
	}

	events, _ := EventService.GetAll()

	w.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(events)
	assertMarshalingError(w, err)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(result)
	assertResponseError(err)
}

func (*controller) GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	email := r.Context().Value("email")

	data := database.GlobalDB.Where("email = ?", email.(string)).First(&user)

	if data.Error == gorm.ErrRecordNotFound {
		assertGormError(w, `"error":"user not found"`)
	}

	var event entity.Event
	err := json.NewDecoder(r.Body).Decode(&event)

	res, err := EventService.GetOne(event.Id)
	if err != nil {
		fmt.Println(err)
	}

	result, err := json.Marshal(res)
	assertMarshalingError(w, err)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(result)
	assertResponseError(err)
}

func (*controller) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var event entity.Event
	err := json.NewDecoder(r.Body).Decode(&event)

	e, _ := EventService.Add(event)

	result, err := json.Marshal(e)
	assertMarshalingError(w, err)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(result)
	assertResponseError(err)
}

func (*controller) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var event entity.Event
	err := json.NewDecoder(r.Body).Decode(&event)

	e, _ := EventService.Update(event)

	result, err := json.Marshal(e)
	assertMarshalingError(w, err)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(result)
	assertResponseError(err)
}

func (*controller) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var event entity.Event
	err := json.NewDecoder(r.Body).Decode(&event)

	EventService.Delete(event.Id)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(`"result":"ok`))
	assertResponseError(err)
}
