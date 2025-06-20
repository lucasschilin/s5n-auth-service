package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager interface {
	GenerateToken(claims map[string]interface{}) (string, error)
	ValidateToken(token string) (map[string]interface{}, error)
}

type JWTManager struct {
	secretKey string
}

func NewJWT(secretKey string) TokenManager {
	return &JWTManager{
		secretKey: secretKey,
	}
}

func (j *JWTManager) GenerateToken(claims map[string]interface{}) (string, error) {
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

func (j *JWTManager) ValidateToken(token string) (map[string]interface{}, error) {

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
