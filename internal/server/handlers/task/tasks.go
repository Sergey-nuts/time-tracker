package task

import (
	"context"
	"log/slog"
	"net/http"
	"time-tracker/internal/logger/sl"
	"time-tracker/internal/server/middleware/mwpagination"
	"time-tracker/internal/server/response"
	"time-tracker/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// получение списка тасков
func TasksHandler(ctx context.Context, log *slog.Logger, service service.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.task.TasksHandler"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		uuid := chi.URLParam(r, "uuid")
		if uuid != "" {
			// task, err := service.Get(ctx, uuid)
		}

		pageid := r.Context().Value(mwpagination.PageId)
		log.Debug("pageid", slog.Int("page", pageid.(int)))
		tasks, err := service.Tasks(ctx, pageid.(int))
		if err != nil {
			log.Error("failed to get tasks list", sl.Err(err))

			// TO DO: разобраться со статусом ответа... после render возвращается ответ 200
			render.JSON(w, r, response.Error(http.StatusInternalServerError, "failed to get list tasks"))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Debug("TasksHandler", slog.Any("responseWriter", w))
			log.Debug("TasksHandler", slog.Any("request", *r))

			return
		}

		data := map[string]interface{}{}
		data["tasks"] = tasks

		log.Debug("TasksHandler ", slog.Any("tasks:", data))

		render.JSON(w, r, data)
		// json.NewEncoder(w).Encode(data)
	}
}
