package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		slog.Info(
			"HTTP Request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", http.StatusOK,
			"duration", time.Since(start),
			"ip", r.RemoteAddr,
			"user_agent", r.UserAgent(),
			"requst_id", r.Header.Get("X-Request-ID"),
		)
	})
}
