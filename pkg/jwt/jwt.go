package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// ---

const (
	TokenType = "JWT"
	secret    = "9b7e4f37-c86f-456c-ac81-9c9f63a1715d"
)

var DefaultAlgorithm = jwt.SigningMethodHS256

// ---

func buildStandardClaims(timeGap time.Duration, subject string) jwt.StandardClaims {
	now := time.Now()
	nowUnix := now.Unix()

	id := uuid.New().String()

	c := jwt.StandardClaims{
		// ExpiresAt: now.Add(timeGap).Unix(),
		ExpiresAt: nowUnix,
		Id:        id,
		IssuedAt:  nowUnix,
		Issuer:    "I2D",
		Subject:   subject,
	}

	return c
}

func New(timeGap time.Duration, subject string) string {

	claims := buildStandardClaims(timeGap, subject)

	token := jwt.NewWithClaims(DefaultAlgorithm, claims)

	tokenString, _ := token.SignedString([]byte(secret))

	return tokenString
}

func keyFunc(t *jwt.Token) (interface{}, error) {
	return []byte(secret), nil
}

func Verify(tokenString string) (jwt.Claims, error) {

	claims := &jwt.StandardClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, token.Claims.Valid()
	}

	return claims, nil
}
