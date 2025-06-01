package service

import (
	"time"

	"github.com/lucasschilin/schily-users-api/internal/port"
)

func generateAccessToken(jwtPort port.JWT, userID string) (string, error) {
	// expiration := time.Now().Add(15 * time.Minute).Unix()
	expiration := time.Now().Add(15 * time.Second).Unix()
	mapClaims := map[string]interface{}{
		"iat":  time.Now().Unix(),
		"exp":  expiration,
		"sub":  userID,
		"type": "access_token",
	}
	return jwtPort.GenerateToken(mapClaims)
}

func generateRefreshToken(jwtPort port.JWT, userID string) (string, error) {
	// expiration := time.Now().Add(1 * time.Hour).Unix()
	expiration := time.Now().Add(1 * time.Minute).Unix()
	mapClaims := map[string]interface{}{
		"iat":  time.Now().Unix(),
		"exp":  expiration,
		"sub":  userID,
		"type": "refresh_token",
	}
	return jwtPort.GenerateToken(mapClaims)
}
