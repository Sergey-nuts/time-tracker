package service

import (
	"context"
	"time-tracker/internal/server/apimodel"
	"time-tracker/internal/service/servicemodel"
)

type TaskService interface {
	Tasks(ctx context.Context) ([]servicemodel.Task, error)
	Add(ctx context.Context, task servicemodel.Task) (string, error)
	//Get(ctx context.Context, uuid string) (*servicemodel.Task, error)
	Edit(ctx context.Context, task servicemodel.Task) (apimodel.Task, error)
}
