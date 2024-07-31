package main

import (
	"fmt"
	"log/slog"
	"time-tracker/internal/config"
	"time-tracker/internal/logger/sl"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg.Env)

	log := sl.SetupLogger(cfg.Env)

	log.Info(
		"time-tracker",
		slog.String("env", cfg.Env),
		slog.String("http_server", cfg.HTTPServer.Address),
	)
	log.Debug("debug messages are enabled")
}
