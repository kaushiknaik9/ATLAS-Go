package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/kaushiknaik9/ATLAS-Go/internal/config"
)

func Home(w http.ResponseWriter, r *http.Request) {
	slog.Info("Server Running at PORT: ", slog.String("address: ", config.LoadConfig().Host))
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "Server Running",
	})
}
