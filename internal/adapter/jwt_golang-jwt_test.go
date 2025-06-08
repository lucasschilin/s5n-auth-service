package adapter_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/lucasschilin/s5n-auth-service/internal/adapter"
)

func TestGenerationToken(t *testing.T) {

	token := map[string]interface{}{
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(30 * time.Minute).Unix(),
		"sub":  "user_id",
		"type": "access_token",
	}

	t.Run("Generate Valid Token", func(t *testing.T) {
		jwtAdapter := adapter.NewJWT("secret_key")

		token, err := jwtAdapter.GenerateToken(token)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if token == "" {
			t.Errorf("Expected valid token, got empty string")
		}

		tokenRegex := regexp.MustCompile(`^[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+$`)
		if !tokenRegex.MatchString(token) {
			t.Errorf("Expected <HEADER>.<PAYLOAD>.<SIGNATURE>, got wrong format string")
		}
	})
}
