package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time-tracker/internal/config"
	"time-tracker/internal/logger/sl"
	"time-tracker/internal/repository"
	"time-tracker/internal/repository/memdb"
	"time-tracker/internal/server/api"
	"time-tracker/internal/service"
	"time-tracker/internal/service/task"
)

type App struct {
	cfg     config.Config
	log     *slog.Logger
	server  *api.API
	service service.TaskService
	repo    repository.Storage
}

func New(ctx context.Context) *App {
	// config
	cfg := config.MustLoad()

	// logger
	log := initLog(cfg.Env)

	// Repository
	// TO DO: postgresql implementation
	// repo := postgresql.New()

	// memdb
	repo := memdb.New(log)

	// service
	service := initService(repo, log)

	// http server
	server := api.New(ctx, service, log)

	a := &App{
		cfg:     cfg,
		log:     log,
		server:  server,
		service: service,
		repo:    repo,
	}

	return a
}

func (a *App) Run(ctx context.Context) {
	srv := &http.Server{
		Addr:         a.cfg.HTTPServer.Address,
		Handler:      a.server.Router(),
		ReadTimeout:  a.cfg.HTTPServer.Timeout,
		WriteTimeout: a.cfg.HTTPServer.Timeout,
		IdleTimeout:  a.cfg.HTTPServer.IdleTimeout,
	}

	done := make(chan os.Signal, 1)
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// goroutine for graceful shutdown
	go func() {
		defer close(done)

		<-sigterm
		a.log.Info("stopping server")
		ctx, cancel := context.WithTimeout(ctx, a.cfg.HTTPServer.ShutdownTimeout)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			a.log.Error("failed to stop server", sl.Err(err))
		}
	}()

	a.log.Info("starting server", slog.String("address", a.cfg.HTTPServer.Address))

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		signal.Stop(sigterm)
		a.log.Error("failed to start server")
		os.Exit(1)
	}

	<-done
	a.log.Info("server stopped")
}

func (a App) Log() *slog.Logger {
	return a.log
}

func (a App) Cfg() config.Config {
	return a.cfg
}

func initService(repo repository.Storage, log *slog.Logger) service.TaskService {
	return task.NewService(repo, log)
}

func initLog(env string) *slog.Logger {
	log := sl.SetupLogger(env)

	return log
}
