package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Logging the HealthCheck Controller")

	statusCode := strconv.Itoa(http.StatusOK)
	w.Header().Set(
		"Content-Type", "application/json",
	)
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(map[string]string{
		"Status":     "OK",
		"StatusCode": statusCode,
		"Message":    "HealthCheck point Running Good",
	})
	if err != nil {
		slog.Info("Health Route not working something wrong with App", slog.String("Error: %s", err.Error()))
	}
}
