package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtSecrectKey = "authentication"
	defaultTokenDuration = 24 * time.Hour
)

type JWTClaims struct {
	UserId uint `json:"user_id"`
	Name string `json:""`
	Email string `json:""`
	jwt.RegisteredClaims
}

// If any function want to behave of JWTService then him obey this interafce
type JWTService interface {
	GenerateToken(userId uint, name string, email string)(string, error)
	// VerifyToken(tokenstr string)(*JWTClaims, error)
}

type JWT struct {
	// SecretKey []byte
	SecretKey string
	Duration  time.Duration
}

func NewJWT(secret string, duration time.Duration) JWTService{

	// NewJWT inove time if not parse secret key then this statement work
	if secret == "" {
		secret = jwtSecrectKey
	}

	// NewJWT inove time if not parse token duration time then this statement work
	if duration == 0 {
		duration = defaultTokenDuration
	}

	return &JWT{
		SecretKey: secret,
		Duration: duration,
	}
}

func (j *JWT) GenerateToken(userId uint, name string, email string) (string, error){
	claims := JWTClaims{
		UserId: userId,
		Name: name,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.Duration)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer: "gotickets",
		},
	}

	// useing claims generate a token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokeStr, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		return " ", err
	}

	return  tokeStr, nil
}

// func (j *JWT) VarifyToken(tokenstr string) {}