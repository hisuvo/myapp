package users

import (
	"myapp/internal/auth"
	"myapp/internal/middlewares"

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
	api.POST("/login", userHandler.LoginUser, middlewares.CheckMiddleware("suvo datta check"))
	api.GET("/me",userHandler.GetMe, middlewares.AuthMiddleware(newJwt))
}

//* Entity -> Register -> Handler -> Service -> Repository