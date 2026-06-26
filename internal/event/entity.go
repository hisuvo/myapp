package event

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model

	Title       string `gorm:"type:varchar(150);not null" json:"title"`
	Description string `gorm:"type:text;not null" json:"description"`
	Location    string `gorm:"type:varchar(255);not null" json:"location"`

	StartAt time.Time `gorm:"not null;index" json:"start_at"`
	EndAt   time.Time `gorm:"not null" json:"end_at"`

	TotalTickets     int `gorm:"not null" json:"total_tickets"`
	AvailableTickets int `gorm:"not null" json:"available_tickets"`

	Price        float64 `gorm:"type:decimal(10,2);default:0" json:"price"`
	// ImageURL     string `gorm:"type:varchar(500)" json:"image_url"`
	// Category     string `gorm:"type:varchar(100)" json:"category"`
	// Status       string `gorm:"type:varchar(20);default:'upcoming'" json:"status"` // upcoming, ongoing, completed, cancelled
	// OrganizerID  uint `gorm:"index" json:"organizer_id"`
}
