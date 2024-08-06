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

func TaskStartHandler(ctx context.Context, log *slog.Logger, service service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.task.StartHandler"

		log := log.With(
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

		log.Debug("request body decoded", slog.Any("task", task))

		taskStarted, err := service.Start(ctx, converter.TaskToService(task))
		if err != nil {
			log.Error("failed to start task", sl.Err(err))
			http.Error(w, "failed to start task", http.StatusInternalServerError)

			return
		}

		log.Debug("start task complete", slog.Any("uuid", taskStarted.UUID))

		data := map[string]interface{}{}
		data["task"] = taskStarted

		render.JSON(w, r, data)
	}
}
