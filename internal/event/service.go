package event

import (
	"myapp/internal/event/dto"
)

type Service interface {
	Create(req *dto.CreateEventRequest) (*dto.EventResponse, error)
	GetAll() (*[]dto.EventResponse, error)
	GetById(eventId uint) (*dto.EventResponse, error)
	Update(eventId uint, req *dto.UpdateEventRequest) (*dto.EventResponse, error)
	// Delete(event *Event) error
}

type service struct{
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(req *dto.CreateEventRequest)(*dto.EventResponse, error){
	event := &Event{
		Title:              req.Title,
		Description:        req.Description,
		Location:           req.Location,
		StartAt:            req.StartAt,
		EndAt:              req.EndAt,
		TotalTickets:       req.TotalTickets,
		AvailableTickets:   req.AvailableTickets,
	}

	if err := s.repository.Create(event); err != nil {
		return nil, err
	}

	response := ToEventResponse(event)

	return &response, nil
}

func (s *service) GetAll()(*[]dto.EventResponse, error){
	events, err := s.repository.GetAll()

	if err !=nil {
		return nil, err
	}

	response := ToEventResponses(*events)
	return &response, nil
}

func (s *service) GetById(eventId uint) (*dto.EventResponse, error) {
	event, err := s.repository.GetById(eventId)

	if err != nil {
		return nil, err
	}

	response := ToEventResponse(event)
	return &response, nil
}

func (s *service) Update(eventId uint,req *dto.UpdateEventRequest)(*dto.EventResponse, error) {
	event, err := s.repository.GetById(eventId)

	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		event.Title = *req.Title
	}

	if req.Description != nil {
		event.Description = *req.Description
	}

	if req.Location != nil {
		event.Location = *req.Location
	}

	if req.StartAt != nil {
		event.StartAt = *req.StartAt
	}

	if req.EndAt != nil {
		event.EndAt = *req.EndAt
	}

	if req.TotalTickets != nil {
		event.TotalTickets = *req.TotalTickets
	}

	if req.AvailableTickets != nil {
		event.AvailableTickets = *req.AvailableTickets
	}

	if err := s.repository.Update(event); err != nil {
		return nil, err
	}

	response := ToEventResponse(event)

	return &response, nil
}



// Todo:
/*
 ?  I recommend one small improvement
 *  
 *  Instead of:
 *  
 *  GetAll() (*[]Event, error)
 *  
 *  return:
 *  
 *  GetAll() ([]Event, error)
 *  
 *  Slices in Go are already reference-like types. Returning a pointer to a slice (*[]Event) is rarely necessary and makes the code more cumbersome.
 *  
 *  So your repository would become:
 *  
 *  type Repository interface {
 *  	Create(event *Event) error
 *  	GetAll() ([]Event, error)
 *  	GetById(eventId uint) (*Event, error)
 *  	Update(event *Event) error
 *  	Delete(event *Event) error
 *  }
 *  
 *  This is the style you'll see in most Go projects because it's simpler and more idiomatic.
 *  
*/