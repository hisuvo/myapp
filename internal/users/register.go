package users

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, db *gorm.DB){
	userRepository := NewRepository(db)
	userService := NewService(userRepository)
	userHandler := NewHandler(userService)

	api := e.Group("/api/v1")
	api.POST("/users", userHandler.CreateUser)
}