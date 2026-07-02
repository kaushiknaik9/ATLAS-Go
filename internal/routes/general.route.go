package routes

import (
	"net/http"

	"github.com/kaushiknaik9/ATLAS-Go/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("GET /", handlers.Home)

	mux.HandleFunc("GET /health", handlers.HealthHandler)

	mux.HandleFunc("GET /about", handlers.AboutPage())

}
