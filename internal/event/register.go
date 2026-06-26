package event

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	enventRegister := NewRepository(db)
	eventService := NewService(enventRegister)
	eventHandler := NewHandler(eventService)

	events := e.Group("/api/v1")

	events.POST("/events", eventHandler.Create)
	events.GET("/events", eventHandler.GetAll)
	events.GET("/events/:id", eventHandler.GetById)
	events.PUT("/events/:id", eventHandler.Update)
}