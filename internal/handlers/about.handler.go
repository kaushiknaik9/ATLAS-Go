package handlers

import (
	"net/http"

	"github.com/kaushiknaik9/ATLAS-Go/internal/response"
)

func AboutPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.SendResponse(w, http.StatusOK, map[string]string{
			"status":  "ok",
			"message": "This is AboutPage of Atlas",
		})
	}
}
