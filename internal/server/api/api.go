package api

import (
	"context"
	"log/slog"
	"net/http"
	"time-tracker/internal/server/handlers/task"
	"time-tracker/internal/server/middleware/mwlogger"
	"time-tracker/internal/server/middleware/mwpagination"
	"time-tracker/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type API struct {
	r *chi.Mux // router
	//repo    service.TaskService // database
	service service.TaskService // service
	log     *slog.Logger        // application logger
}

func New(ctx context.Context, service service.TaskService, log *slog.Logger) *API {
	api := API{
		r:       chi.NewRouter(),
		service: service,
		log:     log,
	}

	api.endpoints(ctx)

	return &api
}

func (a *API) endpoints(ctx context.Context) {
	a.r.Use(middleware.RequestID)
	a.r.Use(mwlogger.New(a.log))
	a.r.Use(middleware.Recoverer)
	a.r.Use(middleware.URLFormat)

	a.r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	a.r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	a.r.Route("/tasks", func(r chi.Router) {
		r.With(mwpagination.Paginate).Get("/", task.TasksHandler(ctx, a.log, a.service))
		r.Post("/", task.AddTaskHandler(ctx, a.log, a.service))
		r.Put("/", task.EditTaskHandler(ctx, a.log, a.service))
		r.Delete("/", task.TaskDeleteHandler(ctx, a.log, a.service))

		r.Get("/start", task.TaskStartHandler(ctx, a.log, a.service))
		r.Get("/stop", task.TaskStopHandler(ctx, a.log, a.service))
	})
}

func (a *API) Router() *chi.Mux {
	return a.r
}
