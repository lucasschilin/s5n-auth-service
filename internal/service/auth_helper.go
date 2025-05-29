package service

import (
	"time"

	"github.com/lucasschilin/schily-users-api/internal/port"
)

func generateAccessToken(jwtPort port.JWT, userID string) (string, error) {
	accessTokenExpiration := time.Now().Add(30 * time.Minute).Unix()
	mapClaims := map[string]interface{}{
		"iat":  time.Now().Unix(),
		"exp":  accessTokenExpiration,
		"sub":  userID,
		"type": "access_token",
	}
	return jwtPort.GenerateToken(mapClaims)
}

func generateRefreshToken(jwtPort port.JWT, userID string) (string, error) {
	accessTokenExpiration := time.Now().Add(30 * time.Minute).Unix()
	mapClaims := map[string]interface{}{
		"iat":  time.Now().Unix(),
		"exp":  accessTokenExpiration,
		"sub":  userID,
		"type": "access_token",
	}
	return jwtPort.GenerateToken(mapClaims)
}
