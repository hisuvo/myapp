package middlewares

import (
	"fmt"
	"myapp/internal/auth"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
)

func AuthMiddleware(JWTService auth.JWTService) echo.MiddlewareFunc{
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		
		return func(c *echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized,"authorization header required")
			}

			parts := strings.Split(authHeader, " ")

			if len(parts) != 2 || parts[0] != "Bearer"{
				return c.JSON(http.StatusUnauthorized,"invalid authorization format")
			}

			tokenString := parts[1]

			claims, err := JWTService.VerifyToken(tokenString)

			if err != nil {
				return c.JSON(http.StatusUnauthorized,"authorization header required")
			}

			c.Set("id", claims.ID)
			c.Set("name", claims.Name)
			c.Set("email", claims.Email)

		 return next(c)
		}
	}
}

func CheckMiddleware(name string) echo.MiddlewareFunc{
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			fmt.Println("check middlerare", name)
			return next(c)
		}
	}

}