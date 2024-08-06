package repository

import (
	"context"
	"errors"
	"time-tracker/internal/repository/repomodel"
	"time-tracker/internal/service/servicemodel"
)

var (
	ErrorTaskNotFound = errors.New("task not found")
)

type Storage interface {
	Tasks(ctx context.Context) ([]servicemodel.Task, error)
	AddTask(ctx context.Context, task repomodel.Task) (servicemodel.Task, error)
	EditTask(ctx context.Context, task repomodel.Task) (servicemodel.Task, error)
	DeleteTask(ctx context.Context, task repomodel.Task) (servicemodel.Task, error)
	Get(ctx context.Context, uuid string) (servicemodel.Task, error)
	// StopTask(ctx context.Context, Task) error
}
