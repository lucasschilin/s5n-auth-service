package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/lucasschilin/s5n-auth-service/internal/dto"
)

func CheckAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		if accessToken == "" || !strings.HasPrefix(accessToken, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(dto.DefaultDetailResponse{
				Detail: "Invalid or missing access token",
			})
			return
		}

		//TODO: Continue the access token validation

		next.ServeHTTP(w, r)
	})
}
