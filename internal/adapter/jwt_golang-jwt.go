package adapter

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/lucasschilin/schily-users-api/internal/port"
)

type JWT struct {
	secretKey string
}

func NewJWT(secretKey string) port.JWT {
	return &JWT{
		secretKey: secretKey,
	}
}

func (j *JWT) GenerateToken(claims map[string]interface{}) (string, error) {
	mapClaims := jwt.MapClaims{}
	for key, value := range claims {
		mapClaims[key] = value
	}

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	token, err := tokenObj.SignedString([]byte(j.secretKey))

	if err != nil {
		return "", err
	}

	return token, nil
}
