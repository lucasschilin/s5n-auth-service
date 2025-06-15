package middleware

import (
	"fmt"
	"net/http"

	"github.com/lucasschilin/s5n-auth-service/internal/cache"
)

func RateLimit(cache cache.Cache) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Rate Limiting Middleware has been actioned")

			next.ServeHTTP(w, r)
		})
	}

}
