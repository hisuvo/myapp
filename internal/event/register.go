package event

import (
	"github.com/labstack/echo/v5"
)

func RegisterRoutes(g *echo.Group, handler *Handler) {
	events := g.Group("/events")

	events.POST("", handler.Create)
	events.GET("", handler.GetAll)
	events.GET("/:id", handler.GetById)
	events.PUT("/:id", handler.Update)
}