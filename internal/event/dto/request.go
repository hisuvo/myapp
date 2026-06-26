package dto

import "time"

type CreateEventRequest struct {
	Title            string 	`json:"title" validate:"required,min=3,max=150"`
	Description      string 	`json:"description" validate:"required,min=10"`
	Location         string 	`json:"location" validate:"required,max=255"`
	StartAt          time.Time 	`json:"start_at" validate:"required"`
	EndAt            time.Time 	`json:"end_at" validate:"required"`
	TotalTickets     int    	`json:"total_tickets" validate:"required,min=1"`
	AvailableTickets int    	`json:"available_tickets" validate:"required,min=0"`
}

type UpdateEventRequest struct {
	Title            *string 	`json:"title" validate:"omitempty,min=3,max=150"`
	Description      *string 	`json:"description" validate:"omitempty,min=10"`
	Location         *string 	`json:"location" validate:"omitempty,max=255"`
	StartAt          *time.Time `json:"start_at" validate:"omitempty"`
	EndAt            *time.Time `json:"end_at" validate:"omitempty"`
	TotalTickets     *int    	`json:"total_tickets" validate:"omitempty,min=1"`
	AvailableTickets *int    	`json:"available_tickets" validate:"omitempty,min=0"`
}

/*
 ?	Why use pointers (*string, *int)
 *
 Todo: Suppose the client sends
 *
 *	{
 *		"title": "Go Conference 2026"
 *	}
 *
 Todo:	Then
 *
 *	Title != nil → Update title
 *	Description == nil → Don't update description
 *	Location == nil → Don't update location
 *
 *	Without pointers, Go cannot distinguish between:
 *
 *	Field not sent
 *	Field sent with a zero value

*/

/*
{
    "title": "Go Conference 2026",
    "description": "The largest Golang conference in Bangladesh.",
    "location": "Dhaka",
    "start_at": "2026-07-15T09:00:00Z",
    "end_at": "2026-07-15T18:00:00Z",
    "total_tickets": 500,
    "available_tickets": 500
}
*/