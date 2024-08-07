package task

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
	"time-tracker/internal/config"
	"time-tracker/internal/logger/sl"
	"time-tracker/internal/repository"
	"time-tracker/internal/server/apimodel"
	def "time-tracker/internal/service"
	"time-tracker/internal/service/converter"
	"time-tracker/internal/service/servicemodel"
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

func (s *service) Tasks(ctx context.Context, pageid int) ([]servicemodel.Task, error) {
	const op = "service.task.Tasks"

	log := s.log.With(
		slog.String("op", op),
	)

	// TO DO
	ctx, cancel := context.WithTimeout(ctx, config.TimeOut)
	defer cancel()

	log.Debug("Tasks", slog.Any("get tasks", "from repository"))

	tasks, err := s.taskRepository.Tasks(ctx, pageid)
	if err != nil {
		log.Error("failed to get Tasks", sl.Err(err))

		return nil, err
	}

	log.Debug("seccess to get Tasks")

	return tasks, nil
}

func (s *service) Add(ctx context.Context, task servicemodel.Task) (apimodel.Task, error) {
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
	task.Running = false
	task.Duration = time.Duration(0)
	task.StartTime = time.Time{}

	task, err := s.taskRepository.AddTask(ctx, converter.TaskToRepo(task))
	if err != nil {
		log.Error("failed to add Task to repo", sl.Err(err))

		return apimodel.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	log.Debug("add task succeeded", slog.Any("added task", task))
	return converter.TaskToApi(task), nil
}

func (s *service) Get(ctx context.Context, uuid string) (apimodel.Task, error) {
	const op = "service.task.Get"

	log := s.log.With(
		slog.String("op", op),
	)

	// TO DO

	ctx, cancel := context.WithTimeout(ctx, config.TimeOut)
	defer cancel()

	task, err := s.taskRepository.Get(ctx, uuid)
	if err != nil {
		log.Error("failed to get Task in repo", sl.Err(err))

		return apimodel.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	return converter.TaskToApi(task), nil
}

func (s *service) Edit(ctx context.Context, task servicemodel.Task) (apimodel.Task, error) {
	const op = "service.task.Edit"

	log := s.log.With(
		slog.String("op", op),
	)

	ctx, cancel := context.WithTimeout(ctx, config.TimeOut)
	defer cancel()

	edited, err := s.taskRepository.EditTask(ctx, converter.TaskToRepo(task))
	if err != nil {
		log.Error("failed to edit Task in repo", sl.Err(err))

		return apimodel.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	return converter.TaskToApi(edited), nil
}

func (s *service) Delete(ctx context.Context, task servicemodel.Task) (apimodel.Task, error) {
	const op = "service.task.Delete"

	log := s.log.With(
		slog.String("op", op),
	)

	ctx, cancel := context.WithTimeout(ctx, config.TimeOut)
	defer cancel()

	deleted, err := s.taskRepository.DeleteTask(ctx, converter.TaskToRepo(task))
	if err != nil {
		log.Error("failed to delete Task in repo", sl.Err(err))

		return apimodel.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	return converter.TaskToApi(deleted), nil
}

func (s *service) Start(ctx context.Context, task servicemodel.Task) (apimodel.Task, error) {
	const op = "service.task.Start"

	log := s.log.With(
		slog.String("op", op),
	)

	ctx, cancel := context.WithTimeout(ctx, config.TimeOut)
	defer cancel()

	t, err := s.taskRepository.Get(ctx, task.UUID)
	if errors.Is(err, repository.ErrorTaskNotFound) {
		log.Error("faild to start Task", sl.Err(err))

		return apimodel.Task{}, err
	}

	if err != nil {
		log.Error("failed to start Task", sl.Err(err))

		return apimodel.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	if t.Running {
		log.Error("failed to start Task", sl.Err(def.ErrorTaskRunning))

		return apimodel.Task{}, def.ErrorTaskRunning
	}

	t.Running = true
	t.StartTime = time.Now()

	t, err = s.taskRepository.EditTask(ctx, converter.TaskToRepo(t))
	if err != nil {
		log.Error("failed to start Task in repo", sl.Err(err))

		return apimodel.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	return converter.TaskToApi(t), nil
}

func (s *service) Stop(ctx context.Context, task servicemodel.Task) (apimodel.Task, error) {
	const op = "service.task.Stop"

	log := s.log.With(
		slog.String("op", op),
	)

	ctx, cancel := context.WithTimeout(ctx, config.TimeOut)
	defer cancel()

	t, err := s.taskRepository.Get(ctx, task.UUID)
	if errors.Is(err, repository.ErrorTaskNotFound) {
		log.Error("faild to stop Task", sl.Err(err))

		return apimodel.Task{}, err
	}

	if err != nil {
		log.Error("failed to stop Task", sl.Err(err))

		return apimodel.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	if !t.Running {
		log.Error("failed to stop Task", sl.Err(def.ErrorTaskNotRunning))

		return apimodel.Task{}, def.ErrorTaskNotRunning
	}

	// TO DO: implement stopping logic
	t.Running = false

	t.Duration = t.Duration + time.Since(t.StartTime)
	t.StartTime = time.Time{}

	t, err = s.taskRepository.EditTask(ctx, converter.TaskToRepo(t))
	if err != nil {
		log.Error("failed to stop Task in repo", sl.Err(err))

		return apimodel.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	return converter.TaskToApi(t), nil
}
