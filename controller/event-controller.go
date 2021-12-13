package controller

import (
	"encoding/json"
	"fmt"
	"log"
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

type Event struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DateTime    string `json:"dateTime"`
	Duration    string `json:"duration"`
	Notes       string `json:"notes"`
}

var events []Event

func init() {
	events = []Event{
		{"1", "Title1", "Description1", "DateTiem", "Duration1", "Notes1"},
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

func assertResponseError(err error) {
	if err != nil {
		return
	}
}

func (*controller) GetAll(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(events)
	assertMarshalingError(w, err)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(result)
	assertResponseError(err)
}

func (*controller) GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)

	var res interface{}

	for _, e := range events {
		if e.Id == event.Id {
			res = e
		}
	}

	result, err := json.Marshal(res)
	assertMarshalingError(w, err)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(result)
	assertResponseError(err)
}

func (*controller) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)

	events = append(events, event)

	result, err := json.Marshal(event)
	assertMarshalingError(w, err)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(result)
	assertResponseError(err)
}

func (*controller) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)

	var pos int

	for i, e := range events {
		if e.Id == event.Id {
			pos = i
		}
	}

	events[pos].Title = event.Title
	events[pos].Description = event.Description
	events[pos].Duration = event.Duration
	events[pos].DateTime = event.DateTime
	events[pos].Notes = event.Notes

	result, err := json.Marshal(events[pos])
	assertMarshalingError(w, err)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(result)
	assertResponseError(err)
}

func removeIndex(s []Event, i int) []Event {
	return append(s[:i], s[i+1:]...)
}

func (*controller) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)

	log.Println(event.Id)

	var pos int
	var exists = false

	for i, e := range events {
		if e.Id == event.Id {
			fmt.Println("Found it")
			exists = true
			pos = i
			break
		} else {
			exists = false
			break
		}
	}

	fmt.Println(pos, events[pos])

	if exists == true {
		events = removeIndex(events, pos)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(`"result":"ok`))
	assertResponseError(err)
}
