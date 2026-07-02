package main

import (
	"log/slog"

	"github.com/kaushiknaik9/ATLAS-Go/internal/config"
	"github.com/kaushiknaik9/ATLAS-Go/internal/server"
)

func main() {
	//Initialze the Config file
	cfg := config.LoadConfig()

	slog.Info("Config file Imported and Readed ", "env", cfg.Env)

	//Setup DataBase

	//Setup Router

	//Setup Server
	server := server.New(cfg)
	server.ListenAndServe()

}
