package middleware

import (
	"net/http"

	"golang.org/x/time/rate"
)

func RateLimiter(limit int) Middleware {
	return func(next http.Handler) http.Handler {
		// Limit requests
		limiter := rate.NewLimiter(1, limit)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				w.WriteHeader(http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

}
