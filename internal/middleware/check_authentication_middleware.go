package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
	"github.com/lucasschilin/s5n-auth-service/internal/port"
)

type contextKey string

const UserIDKey = contextKey("user_id")

func CheckAuthentication(jwtPort port.JWT) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			res := dto.DefaultDetailResponse{
				Detail: "Invalid or missing access token",
			}

			accessToken := r.Header.Get("Authorization")
			if accessToken == "" || !strings.HasPrefix(accessToken, "Bearer ") {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(res)
				return
			}

			accessTokenString := strings.TrimPrefix(accessToken, "Bearer ")

			accessTokenClaims, err := jwtPort.ValidateToken(accessTokenString)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(res)
				return
			}

			typeClaim, exists := accessTokenClaims["type"]
			if !exists {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(res)
				return
			}
			if typeClaim, ok := typeClaim.(string); !ok || typeClaim != "access_token" {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(res)
				return
			}

			sub, exists := accessTokenClaims["sub"]
			if !exists {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(res)
				return
			}
			userID, ok := sub.(string)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(res)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
