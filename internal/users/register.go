package users

import (
	"myapp/internal/auth"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, db *gorm.DB){
	userRepository := NewRepository(db)
	newJwt := auth.NewJWT("",0)
	userService := NewService(userRepository, newJwt)
	userHandler := NewHandler(userService)

	api := e.Group("/api/v1/auth")
	api.POST("/register", userHandler.CreateUser)
	api.POST("/login", userHandler.LoginUser)
}

//* Entity -> Register -> Handler -> Service -> Repository