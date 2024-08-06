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

func TaskStopHandler(ctx context.Context, log *slog.Logger, service service.TaskService) http.HandlerFunc {
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

		taskStopped, err := service.Stop(ctx, converter.TaskToService(task))
		if err != nil {
			log.Error("failed to stop task", sl.Err(err))
			http.Error(w, "failed to stop task", http.StatusInternalServerError)

			return
		}

		log.Debug("stop task complete", slog.Any("uuid", taskStopped.UUID))

		data := map[string]interface{}{}
		data["task"] = taskStopped

		render.JSON(w, r, data)
	}
}
