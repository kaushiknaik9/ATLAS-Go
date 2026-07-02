package routes

import (
	"net/http"

	"github.com/kaushiknaik9/ATLAS-Go/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/", handlers.Home)

	mux.HandleFunc("/health", handlers.HealthHandler)

}
