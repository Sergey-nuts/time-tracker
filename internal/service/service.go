package service

import (
	"context"
	"errors"
	"time-tracker/internal/server/apimodel"
	"time-tracker/internal/service/servicemodel"
)

var (
	ErrorTaskRunning    = errors.New("task alredy running")
	ErrorTaskNotRunning = errors.New("task not running")
)

type TaskService interface {
	Tasks(ctx context.Context, pageid int) ([]servicemodel.Task, error)
	Add(ctx context.Context, task servicemodel.Task) (apimodel.Task, error)
	Get(ctx context.Context, uuid string) (apimodel.Task, error)
	Edit(ctx context.Context, task servicemodel.Task) (apimodel.Task, error)
	Delete(ctx context.Context, task servicemodel.Task) (apimodel.Task, error)
	Start(ctx context.Context, task servicemodel.Task) (apimodel.Task, error)
	Stop(ctx context.Context, task servicemodel.Task) (apimodel.Task, error)
}
