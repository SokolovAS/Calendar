package grps

import (
	"Calendar/entity"
	"Calendar/internal/server/http"
	"Calendar/internal/services/calendar"
	_ "Calendar/internal/services/calendar"
	"Calendar/pb"
	"context"
	_ "errors"
	_ "go/types"
	_ "log"
)

type Server struct {
	eventServ http.EventService
}

func NewGRPCServer(e *calendar.EventService) *Server {
	return &Server{
		eventServ: e,
	}
}

func (s Server) GetAll(context.Context, *pb.Event) (*pb.EventsResponse, error) {
	events, err := s.eventServ.GetAll()
	if err != nil {
		return nil, err
	}
	var res []*pb.Event
	for _, e := range events {
		eProto := e.ToProto()
		res = append(res, eProto)
	}
	return &pb.EventsResponse{
		Events: res,
	}, nil
}

func (s Server) GetOne(ctx context.Context, event *pb.Event) (*pb.Event, error) {
	e := entity.FromProto(event)
	res, err := s.eventServ.GetOne(e.Id)
	if err != nil {
		return &pb.Event{}, err
	}
	return res.ToProto(), err
}

func (s Server) Add(ctx context.Context, event *pb.Event) (*pb.Event, error) {
	e := entity.FromProto(event)
	res, err := s.eventServ.Add(e)
	if err != nil {
		return &pb.Event{}, err
	}
	return res.ToProto(), err
}

func (s Server) Update(ctx context.Context, event *pb.Event) (*pb.Event, error) {
	e := entity.FromProto(event)
	res, err := s.eventServ.Update(e)
	if err != nil {
		return &pb.Event{}, err
	}
	return res.ToProto(), err
}

func (s Server) Delete(ctx context.Context, event *pb.Event) (*pb.Empty, error) {
	e := entity.FromProto(event)
	err := s.eventServ.Delete(e.Id)
	if err != nil {
		return &pb.Empty{}, err
	}
	return &pb.Empty{}, err
}

func (s Server) MustEmbedUnimplementedEventServiceServer() {
	panic("implement me")
}
