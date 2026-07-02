package main

import (
	"log/slog"

	"github.com/kaushiknaik9/ATLAS-Go/internal/config"
)

func main() {
	//Initialze the Config file
	cfg := config.ConfigImported()

	slog.Info("Config file Imported and Readed ", slog.String("Config file is Read and Env is seton: %s", cfg.Env))
	//Setup DataBase
	//Setup Router
	//Setup Server
}
