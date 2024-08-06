package task

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"time-tracker/internal/logger/sl"
	"time-tracker/internal/server/apimodel"
	"time-tracker/internal/server/converter"
	"time-tracker/internal/service"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// добавление Таски
func AddTaskHandler(ctx context.Context, log *slog.Logger, service service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.task.AddTaskHandler"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var task apimodel.Task
		err := render.DecodeJSON(r.Body, &task)
		if errors.Is(err, io.EOF) {
			log.Error("request body is empty")
			http.Error(w, "empty request body", http.StatusBadRequest)

			return
		}
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			http.Error(w, "failed to decode request", http.StatusInternalServerError)

			return
		}

		log.Info("request body decoded", slog.Any("task", task))

		task, err = service.Add(ctx, converter.TaskToService(task))
		if err != nil {
			log.Error("failed to add task", sl.Err(err))
			http.Error(w, "failed to add task", http.StatusInternalServerError)

			return
		}

		log.Info("add task complete", slog.Any("task", task))

		data := map[string]interface{}{}
		data["task"] = task

		render.JSON(w, r, data)
	}
}
