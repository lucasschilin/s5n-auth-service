package authservice

import (
	"time"

	"github.com/lucasschilin/s5n-auth-service/internal/port"
)

func generateAccessToken(jwtPort port.JWT, userID string) (string, error) {
	expiration := time.Now().Add(15 * time.Minute).Unix()
	return generateToken(jwtPort, "access_token", int(expiration), userID)
}

func generateRefreshToken(jwtPort port.JWT, userID string) (string, error) {
	expiration := time.Now().Add(90 * time.Minute).Unix()
	return generateToken(jwtPort, "refresh_token", int(expiration), userID)
}

func generateToken(
	jwtPort port.JWT, tokenType string, expiration int, userID string,
) (string, error) {
	mapClaims := map[string]interface{}{
		"iat":  time.Now().Unix(),
		"exp":  expiration,
		"sub":  userID,
		"type": tokenType,
	}
	return jwtPort.GenerateToken(mapClaims)
}
