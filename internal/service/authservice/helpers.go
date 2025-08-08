package authservice

import (
	"time"

	"github.com/lucasschilin/s5n-auth-service/internal/service/authservice/jwt"
)

func generateAccessToken(tokenManager jwt.TokenManager, userID string) (string, error) {
	expiration := time.Now().Add(15 * time.Minute).Unix()
	return generateToken(tokenManager, "access_token", int(expiration), userID)
}

func generateRefreshToken(tokenManager jwt.TokenManager, userID string) (string, error) {
	expiration := time.Now().Add(90 * time.Minute).Unix()
	return generateToken(tokenManager, "refresh_token", int(expiration), userID)
}

func generateToken(
	tokenManager jwt.TokenManager, tokenType string, expiration int, userID string,
) (string, error) {
	mapClaims := map[string]interface{}{
		"iat":  time.Now().Unix(),
		"exp":  expiration,
		"sub":  userID,
		"type": tokenType,
	}
	return tokenManager.GenerateToken(mapClaims)
}
