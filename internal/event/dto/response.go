package dto

import (
	"time"
)

type EventResponse struct {
	ID                 uint      `json:"id"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	Location           string    `json:"location"`
	StartAt            time.Time `json:"start_at"`
	EndAt              time.Time `json:"end_at"`
	TotalTickets       int       `json:"total_tickets"`
	AvailableTickets   int       `json:"available_tickets"`
	BookedTickets      int       `json:"booked_tickets"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
