package adapter

// TODO: refatorar para o formato do cache

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lucasschilin/s5n-auth-service/internal/port"
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

func (j *JWT) ValidateToken(token string) (map[string]interface{}, error) {

	claims := jwt.MapClaims{}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	}

	parsedToken, err := jwt.ParseWithClaims(token, claims, keyFunc)
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil

}
