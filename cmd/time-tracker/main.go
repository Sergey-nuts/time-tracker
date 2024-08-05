package main

import (
	"context"
	"log/slog"
	"time-tracker/internal/pkg/app"
)

func main() {
	ctx := context.Background()
	app := app.New(ctx)

	app.Log().Info(
		"starting time-tracker",
		slog.String("env", app.Cfg().Env),
		slog.String("version", "0.0.1"),
	)
	app.Log().Debug("debug messages are enabled")

	app.Run(ctx)
}
