package calendar

import (
	"Calendar/entity"
	"Calendar/internal/repository"
	"errors"
	"fmt"
)

type RepoPG interface {
	GetAllEvents() ([]entity.Event, error)
	GetOne(id string) (entity.Event, error)
	Add(e entity.Event) error
	Update(e entity.Event) error
	Delete(id string) error
}

type EventService struct {
	repoPG RepoPG
}

func NewEventService(r RepoPG) *EventService {

	return &EventService{
		repoPG: r,
	}
}

type ServiceErr struct {
	Code    int
	Message string
}

func (e ServiceErr) Error() string {
	return fmt.Sprintf("Code %d, message: %v", e.Code, e.Message)
}

func (eS *EventService) GetAll() ([]entity.Event, error) {
	events, err := eS.repoPG.GetAllEvents()
	if err != nil {
		repoErr := repository.RepoError{}
		if errors.As(err, &repoErr) {
			err := ServiceErr{
				Code:    500,
				Message: fmt.Sprintf("%#v\n", repoErr),
			}
			return nil, err
		}

		return nil, err
	}
	return events, nil
}

func (eS *EventService) GetOne(id string) (entity.Event, error) {
	e, err := eS.repoPG.GetOne(id)
	if err != nil {
		repoErr := repository.RepoError{}
		errors.As(err, &repoErr)
		err := ServiceErr{
			Code:    500,
			Message: fmt.Sprintf("%#v\n", repoErr),
		}
		return entity.Event{}, err
	}
	return e, nil
}

func (eS *EventService) Add(event entity.Event) (entity.Event, error) {
	err := eS.repoPG.Add(event)
	if err != nil {
		repoErr := repository.RepoError{}
		errors.As(err, &repoErr)
		err := ServiceErr{
			Code:    500,
			Message: fmt.Sprintf("%#v\n", repoErr),
		}
		return entity.Event{}, err
	}
	return event, nil
}

func (eS *EventService) Update(event entity.Event) (entity.Event, error) {
	err := eS.repoPG.Update(event)
	if err != nil {
		repoErr := repository.RepoError{}
		if errors.As(err, &repoErr) {
			err = ServiceErr{
				Code:    repoErr.Code,
				Message: fmt.Sprintf("%#v\n", repoErr),
			}
		}
		return entity.Event{}, err
	}
	return event, err
}

func (eS *EventService) Delete(id string) error {
	err := eS.repoPG.Delete(id)
	if err != nil {
		repoErr := repository.RepoError{}
		errors.As(err, &repoErr)
		err := ServiceErr{
			Code:    500,
			Message: fmt.Sprintf("%#v\n", repoErr),
		}
		return err
	}
	return err
}
