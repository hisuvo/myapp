একটি Production-Level JWT package সাধারণত token generate, validate, parse এবং config management আলাদা রাখে।

### internal/auth/jwt.go

```go
package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	SecretKey []byte
	Duration  time.Duration
}

type Claims struct {
	UserID uint `json:"user_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func NewJWT(secret string, duration time.Duration) *JWT {
	return &JWT{
		SecretKey: []byte(secret),
		Duration:  duration,
	}
}

// Generate Token
func (j *JWT) GenerateToken(userID uint) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(j.Duration),
			),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.SecretKey)
}

// Validate Token
func (j *JWT) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}

			return j.SecretKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

```

---

### config.yaml

```yaml
jwt:
  secret: "super-secret-key"
  duration: 24h
```

---

### config struct

```go
type JWTConfig struct {
	Secret   string
	Duration time.Duration
}
```

---

### Initialize

```go
jwtService := auth.NewJWT(
	cfg.JWT.Secret,
	cfg.JWT.Duration,
)
```

---

### Generate Token

```go
token, err := jwtService.GenerateToken(user.ID)
if err != nil {
	return err
}
```

---

### Middleware Usage

```go
claims, err := jwtService.ValidateToken(tokenString)
if err != nil {
	return echo.NewHTTPError(
		http.StatusUnauthorized,
		"invalid token",
	)
}

fmt.Println(claims.UserID)
```

### Production Improvements

Production project-এ সাধারণত আরও থাকে:

```text
internal/
└── auth/
    ├── jwt.go
    ├── middleware.go
    ├── claims.go
    └── context.go
```

- `jwt.go` → Generate & Validate Token
- `middleware.go` → Authentication Middleware
- `claims.go` → Custom Claims
- `context.go` → User ID Context-এ Store করা

এই structure Echo, Gin, Fiber, Chi সব Go framework-এ commonly ব্যবহার করা হয়।
