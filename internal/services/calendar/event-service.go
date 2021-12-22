package calendar

import (
	"Calendar/entity"
	"errors"
	"fmt"
)

var events []entity.Event

func init() {
	events = []entity.Event{
		{"1", "Title1", "Description1", "DateTiem", "Duration1", "Notes1"},
	}
}

type EventService struct {
	conn SqliteRepo
}

func NewEventService(r SqliteRepo) *EventService {

	return &EventService{
		conn: r,
	}
}

func (*EventService) GetAll() ([]entity.Event, error) {
	return events, nil
}

func (*EventService) GetOne(id string) (entity.Event, error) {
	var res entity.Event
	var exist bool

	for _, e := range events {
		if e.Id == id {
			res = e
			exist = true
		}
	}

	if exist {
		return res, nil
	}

	return res, errors.New("not able to find the event")
}

func (*EventService) Add(event entity.Event) (entity.Event, error) {
	events = append(events, event)
	return event, nil
}

func (*EventService) Update(event entity.Event) (entity.Event, error) {
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

	return events[pos], nil
}

func (*EventService) Delete(id string) {
	var pos int
	var exists = false

	for i, e := range events {
		if e.Id == id {
			fmt.Println("Found it")
			exists = true
			pos = i
			break
		} else {
			exists = false
		}
	}

	if exists == true {
		events = removeIndex(events, pos)
	}
}

func removeIndex(s []entity.Event, i int) []entity.Event {
	return append(s[:i], s[i+1:]...)
}
