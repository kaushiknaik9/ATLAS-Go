package server

import (
	"net/http"

	"github.com/kaushiknaik9/ATLAS-Go/internal/config"
	"github.com/kaushiknaik9/ATLAS-Go/internal/routes"
)

func New(cfg *config.Config) *http.Server {

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)
	routes.RegisterStudentRoutes(mux)

	server := &http.Server{
		Addr:    cfg.HttpServer.Host,
		Handler: mux,
	}
	return server
}
