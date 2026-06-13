package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtSecrectKey = "authentication"
	defaultTokenDuration = 24 * time.Hour
)

type JWTClaims struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// If any function want to behave of JWTService then him obey this interafce
type JWTService interface {
	GenerateToken(userId uint, name string, email string)(string, error)
	VerifyToken(tokenstr string)(*JWTClaims, error)
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

func (j *JWT) GenerateToken(id uint, name string, email string) (string, error){
	claims := JWTClaims{
		Id: id,
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

func (j *JWT) VerifyToken(tokenstr string)(*JWTClaims, error) {
	
	token, err := jwt.ParseWithClaims(tokenstr,&JWTClaims{},func(t *jwt.Token) (any, error) {
		if _, Ok := t.Method.(*jwt.SigningMethodHMAC); !Ok {
			return nil, fmt.Errorf("Unexpected sign method %v",t.Header["alg"])
		}
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("Unexpected sign method %v", err)
	}

	if claims, Ok := token.Claims.(*JWTClaims); Ok && token.Valid {
		return claims, nil
	}

	return  nil, fmt.Errorf("Invalid token")
}