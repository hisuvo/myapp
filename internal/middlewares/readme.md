SUVO, Echo framework-এর জন্য একটি production-style JWT Auth Middleware উদাহরণ:

```go
package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

var jwtSecret = []byte("your-secret-key")

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"authorization header required",
			)
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"invalid authorization format",
			)
		}

		tokenString := parts[1]

        // validation token here
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"invalid token",
			)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"invalid claims",
			)
		}

		// Save user data in context
		c.Set("user_id", claims["user_id"])
		c.Set("email", claims["email"])

		return next(c)
	}
}
```

### Route এ ব্যবহার

```go
e := echo.New()

private := e.Group("/api")
private.Use(AuthMiddleware)

private.GET("/profile", func(c echo.Context) error {
	userID := c.Get("user_id")

	return c.JSON(200, map[string]any{
		"user_id": userID,
	})
})
```

### Request Example

```http
GET /api/profile

Authorization: Bearer eyJhbGciOiJIUzI1Ni...
```

### Project Structure

```text
internal/
├── middleware/
│   └── auth.go
├── auth/
│   ├── jwt.go
│   └── service.go
├── user/
│   ├── handler.go
│   ├── service.go
│   └── repository.go
```

Production-এ secret key code-এর মধ্যে hardcode না করে `.env` বা config package থেকে load করা উচিত। এছাড়া token expiration, refresh token, role-based authorization (RBAC) এবং custom claims ব্যবহার করা ভালো।
