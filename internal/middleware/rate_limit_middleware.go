package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/lucasschilin/s5n-auth-service/internal/cache"
	"github.com/lucasschilin/s5n-auth-service/internal/dto"
)

const (
	limit                 = 2
	windowDurationSeconds = 1
)

func RateLimit(cache cache.Cache) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			prefix := "rate_limit"
			remoteAddr := r.RemoteAddr
			ip, _, _ := net.SplitHostPort(remoteAddr)

			userAgent := "unknown"
			if r.UserAgent() != "" {
				userAgent = r.UserAgent()
			}
			key := fmt.Sprintf(
				"%s:%s", prefix, userIdentifierHash(ip+userAgent),
			)

			now := time.Now().Unix()
			startWindow := now - windowDurationSeconds

			var reqs []int64

			val, _ := cache.Get(key)
			if val == "" {
				val = "[]"
			}

			json.Unmarshal([]byte(val), &reqs)
			reqs = append(reqs, now)

			firstReqInWindow := 0

			for i, r := range reqs {
				if r >= startWindow {
					break
				}

				firstReqInWindow = i + 1

			}

			reqs = reqs[firstReqInWindow:]

			reqsBytes, _ := json.Marshal(reqs)

			cache.Set(
				key, string(reqsBytes), (windowDurationSeconds+5)*time.Second,
			)

			if len(reqs) > limit {
				w.WriteHeader(http.StatusTooManyRequests)
				json.NewEncoder(w).Encode(dto.DefaultDetailResponse{
					Detail: "Too many requests.",
				})
				return
			}

			next.ServeHTTP(w, r)
		})
	}

}

func userIdentifierHash(userIdentifier string) string {
	uiBytes := []byte(userIdentifier)

	hash := sha256.Sum256(uiBytes)

	return hex.EncodeToString(hash[:])
}
