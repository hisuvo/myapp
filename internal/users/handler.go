/*
* Handler only handle request and reponse
* payload have exist
* validation is accour currectly
 */

package users

import (
	"errors"
	httpresponse "myapp/internal/httpresponse"
	"myapp/internal/users/dto"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

// Conastant function take service and return handler
func NewHandler(service *service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateUser(c *echo.Context) error {
	var req dto.CreateRequest

	if err := c.Bind(&req); err != nil {
		return  c.JSON(http.StatusBadRequest, httpresponse.NewWithDetails(
			http.StatusBadRequest,
			"Invalid request paylad",
			err.Error(),
		))
	}

	if err := c.Validate(&req); err != nil {
		return  c.JSON(http.StatusBadRequest, httpresponse.NewWithDetails(
			http.StatusBadRequest,
			"Validation failed",
			err.Error(),
		))
	}

  response, err := h.service.CreateUser(req)

  if err != nil {
	if errors.Is(err, ErrAlreadyExist) {
		return  c.JSON(http.StatusConflict, httpresponse.NewWithDetails(
			http.StatusConflict,
			"User already exists!",
			err.Error(),
		))
	}

	return  c.JSON(http.StatusBadRequest, httpresponse.NewWithDetails(
		http.StatusBadRequest,
		"Failed to create user",
		err.Error(),
	))
  }

  return c.JSON(http.StatusCreated, response)

}

func (h *handler) LoginUser(c *echo.Context) error {
	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.NewWithDetails(
			http.StatusBadRequest,
			"Invalide request payload",
			err.Error(),
		))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.NewWithDetails(
			http.StatusBadRequest,
			"Invalide request payload",
			err.Error(),
		))
	}
		
	response, err := h.service.LoginUser(req)

	if err != nil {

		if errors.Is(err, ErrAlreadyExist) {
		return  c.JSON(http.StatusUnauthorized, httpresponse.NewWithDetails(
			http.StatusUnauthorized,
			"Invalide email and password",
			err.Error(),
		))
	}

		return  c.JSON(http.StatusInternalServerError, httpresponse.NewWithDetails(
		http.StatusInternalServerError,
		"Failed to loign user",
		err.Error(),
	))
	}
	
	return c.JSON(http.StatusCreated, response)
}

/*
//* Important Note:

* Handler service নেয় কারণ handler শুধু request handle করে,
* আর business logic service layer-এ থাকে

? Interview answer
* Handler depends on service layer to separate concerns.
* It keeps HTTP logic isolated while service handles business rules,
* making the system modular and testable.
*/