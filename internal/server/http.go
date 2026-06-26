package server

import (
	"myapp/internal/config"
	"myapp/internal/event"
	"myapp/internal/users"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

func Start(db *gorm.DB, cfg *config.Config) {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.Validator = &CustomValidator{validator: validator.New()}

	//? GET: all user api
	e.GET("/",func(c *echo.Context) error {
		return c.String(http.StatusOK,"Hello, world! In Go Language inside!")
	})

	//? POST: Create user api
	users.RegisterRoute(e, db)

	event.RegisterRoutes(e, db)

	port := cfg.Port
	if err := e.Start(":"+port); err != nil{
		e.Logger.Error("failed to server start","error", err)
	}
}