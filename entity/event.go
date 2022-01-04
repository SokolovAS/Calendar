package entity

import "Calendar/pb"

type Event struct {
	Id          string `json:"id"`
	IdUser      string `json:"id_user"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DateTime    string `json:"dateTime"`
	Duration    string `json:"duration"`
	Notes       string `json:"notes"`
}

func (e *Event) ToProto() *pb.Event {
	return &pb.Event{
		Id:          e.Id,
		IdUser:      e.IdUser,
		Title:       e.Title,
		Description: e.Description,
		DateTime:    e.DateTime,
		Duration:    e.Duration,
		Notes:       e.Notes,
	}
}

func FromProto(e *pb.Event) Event {
	return Event{
		Id:          e.Id,
		IdUser:      e.IdUser,
		Title:       e.Title,
		Description: e.Description,
		DateTime:    e.DateTime,
		Duration:    e.Duration,
		Notes:       e.Notes,
	}
}
