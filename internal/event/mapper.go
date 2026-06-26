package event

import (
	"myapp/internal/event/dto"
)

// Mapper Function: [Read 📂mapper folder for knowing indeatils about this]
// Instead of manually creating the response in every handler, create a mapper.

func ToEventResponse(event *Event) dto.EventResponse {
	return dto.EventResponse{
		ID:               event.ID,
		Title:            event.Title,
		Description:      event.Description,
		Location:         event.Location,
		StartAt:          event.StartAt,
		EndAt:            event.EndAt,
		TotalTickets:     event.TotalTickets,
		AvailableTickets: event.AvailableTickets,
		BookedTickets:    event.TotalTickets - event.AvailableTickets,
		CreatedAt:        event.CreatedAt,
		UpdatedAt:        event.UpdatedAt,
	}
}

// For List API

func ToEventResponses(events []Event) []dto.EventResponse {
	responses := make([]dto.EventResponse, len(events))

	for i, e := range events {
		responses[i] = ToEventResponse(&e)
	}

	return responses
}
