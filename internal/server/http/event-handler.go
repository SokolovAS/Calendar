package http

import (
	"Calendar/entity"
	"encoding/json"
	"net/http"
)

type EventHandler struct {
	eServ EventService
	uServ UserService
}

//type EventHandler interface {
//	GetAll(w http.ResponseWriter, r *http.Request)
//	GetOne(w http.ResponseWriter, r *http.Request)
//	Add(w http.ResponseWriter, r *http.Request)
//	Update(w http.ResponseWriter, r *http.Request)
//	Delete(w http.ResponseWriter, r *http.Request)
//}

func NewEventHandler(eS EventService, uS UserService) *EventHandler {
	return &EventHandler{
		eServ: eS,
		uServ: uS,
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

func (eventH *EventHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	events, err := eventH.eServ.GetAll()

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(`{"error":"internal server error"}`))
	}

	result, err := json.Marshal(events)
	assertMarshalingError(w, err)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(result)
	assertResponseError(err)
}

func (eventH *EventHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var event entity.Event
	err := json.NewDecoder(r.Body).Decode(&event)

	res, err := eventH.eServ.GetOne(event.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(`{"error":"internal server error"}`))
	}

	result, err := json.Marshal(res)
	assertMarshalingError(w, err)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(result)
	assertResponseError(err)
}

func (eventH *EventHandler) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var event entity.Event
	err := json.NewDecoder(r.Body).Decode(&event)

	e, err := eventH.eServ.Add(event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(`{"error":"internal server error"}`))
	}

	result, err := json.Marshal(e)
	assertMarshalingError(w, err)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(result)
	assertResponseError(err)
}

func (eventH *EventHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var event entity.Event
	err := json.NewDecoder(r.Body).Decode(&event)

	e, err := eventH.eServ.Update(event)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(`{"error":"internal server error"}`))
	}

	result, err := json.Marshal(e)
	assertMarshalingError(w, err)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(result)
	assertResponseError(err)
}

func (eventH *EventHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var event entity.Event
	err := json.NewDecoder(r.Body).Decode(&event)

	eventH.eServ.Delete(event.Id)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(`"result":"ok`))
	assertResponseError(err)
}
