package event

import (
	"myapp/internal/event/dto"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Post /events
func (h *Handler) Create(c *echo.Context) error {
	var req dto.CreateEventRequest

	// Bind
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
		})
	}

	// Validate
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Validation failed",
			"errors":  err.Error(),
		})
	}

	event, err := h.service.Create(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, event)
}

//Get /events
func (h *Handler) GetAll(c *echo.Context) error {
	event, err := h.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message":"Not found events",
		})
	}

	return c.JSON(http.StatusOK, event)
}

//Get /events/:id
func (h *Handler) GetById(c *echo.Context) error {
	// strconv.ParseUint(string, base, bitSize)
	id, err := strconv.ParseUint(c.Param("id"),10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid event id",
		})
	}

	event, err := h.service.GetById(uint(id))

	if err != nil {
		return c.JSON(http.StatusNotFound,map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, event)
}

// Update /events/:id
func (h *Handler) Update(c *echo.Context) error{
	id, err := strconv.ParseUint(c.Param("id"),10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid event id",
		})
	}

	var req dto.UpdateEventRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,map[string]string{
			"message":"Invalid request body",
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest,map[string]string{
			"message":"Validation failed",
			"error":err.Error(),
		})
	}

	event, err := h.service.Update(uint(id), &req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, event)
}