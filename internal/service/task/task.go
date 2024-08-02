package task

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	"time-tracker/internal/config"
	"time-tracker/internal/logger/sl"
	"time-tracker/internal/repository"
	def "time-tracker/internal/service"
	"time-tracker/internal/service/converter"
	model "time-tracker/internal/service/servicemodel"
)

var _ def.TaskService = (*service)(nil)

type service struct {
	taskRepository repository.Storage
	log            *slog.Logger
}

func NewService(taskRepo repository.Storage, log *slog.Logger) *service {
	return &service{
		taskRepository: taskRepo,
		log:            log,
	}
}

func (s *service) Tasks(ctx context.Context) ([]model.Task, error) {
	const op = "service.task.Tasks"

	log := s.log.With(
		slog.String("op", op),
	)

	// TO DO
	ctx, cancel := context.WithTimeout(ctx, config.TimeOut)
	defer cancel()

	log.Debug("Tasks", slog.Any("get tasks", "to repository"))

	tasks, err := s.taskRepository.Tasks(ctx)
	if err != nil {
		log.Error("failed to get Tasks", sl.Err(err))

		return nil, err
	}

	log.Debug("seccess to get Tasks")

	return tasks, nil
}

func (s *service) Add(ctx context.Context, task model.Task) (string, error) {
	const op = "service.task.Add"

	log := s.log.With(
		slog.String("op", op),
	)

	ctx, cancel := context.WithTimeout(ctx, config.TimeOut)
	defer cancel()

	// log.Debug()
	if task.CreationTime.IsZero() {
		task.CreationTime = time.Now()
	}

	uuid, err := s.taskRepository.AddTask(ctx, converter.TaskToRepoFromService(task))
	if err != nil {
		log.Error("failed to add Task to repo", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Debug("add task succeeded")
	return uuid, nil
}

func (s *service) Get(ctx context.Context, uuid string) (*model.Task, error) {
	// TO DO

	return nil, nil
}
