package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/lucasschilin/s5n-auth-service/internal/cache"
)

const (
	limit                 = 3
	windowDurationSeconds = 10
)

func RateLimit(cache cache.Cache) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			prefix := "ratelimit"
			remoteAddr := r.RemoteAddr
			ip, _, _ := net.SplitHostPort(remoteAddr)

			userAgent := "unknown"
			if r.UserAgent() != "" {
				userAgent = r.UserAgent()
			}

			ui := userIdentifierHash(ip + userAgent)

			key := fmt.Sprintf("%s:%s", prefix, ui)

			now := time.Now().Unix()
			startWindow := now - windowDurationSeconds

			reqs, err := cache.Get(key)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(reqs, startWindow)
			// TODO: Continuar implementação (converter valor para json, verificar length, adicionar unix() time, ...)

			next.ServeHTTP(w, r)
		})
	}

}

func userIdentifierHash(userIdentifier string) string {
	uiBytes := []byte(userIdentifier)

	hash := sha256.Sum256(uiBytes)

	return hex.EncodeToString(hash[:])
}
