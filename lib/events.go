package lib

import (
	"github.com/dgrijalva/jwt-go"
)

type EventsJwtClaims struct {
	Events string `json:"events"`
	jwt.StandardClaims
}

func EventsJwtFrom(payload []byte, secret string) (string, error) {
	claims := EventsJwtClaims{
		string(payload),
		jwt.StandardClaims{
			ExpiresAt: 1700000000,
			Issuer:    "atomic-event-mocks",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, err
}
