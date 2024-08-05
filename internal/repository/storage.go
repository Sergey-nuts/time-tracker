package repository

import (
	"context"
	"time-tracker/internal/repository/repomodel"
	"time-tracker/internal/service/servicemodel"
)

type Storage interface {
	Tasks(ctx context.Context) ([]servicemodel.Task, error)
	AddTask(ctx context.Context, task repomodel.Task) (string, error)
	EditTask(ctx context.Context, task repomodel.Task) (servicemodel.Task, error)
	// DeletTask(ctx context.Context, Task) error
	// StartTask(ctx context.Context, Task) error
	// StopTask(ctx context.Context, Task) error
}
