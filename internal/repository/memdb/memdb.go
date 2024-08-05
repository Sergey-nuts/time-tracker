package memdb

import (
	"context"
	"fmt"
	"log/slog"
	"time-tracker/internal/repository"
	"time-tracker/internal/repository/converter"
	"time-tracker/internal/repository/repomodel"
	"time-tracker/internal/service/servicemodel"

	"github.com/google/uuid"
)

var _ repository.Storage = (*InMemDB)(nil)

type InMemDB struct {
	tasks map[string]repomodel.Task
	log   *slog.Logger
}

func New(log *slog.Logger) *InMemDB {
	t := make(map[string]repomodel.Task)

	return &InMemDB{tasks: t, log: log}
}

func (db *InMemDB) Tasks(_ context.Context) ([]servicemodel.Task, error) {
	const op = "repository.InMemory.Tasks"

	t := make([]servicemodel.Task, len(db.tasks))
	for _, v := range db.tasks {
		t = append(t, converter.ToTaskFromRepo(v))
	}

	return t, nil
}

func (db *InMemDB) AddTask(_ context.Context, task repomodel.Task) (string, error) {
	const op = "repository.InMemory.AddTask"

	log := db.log.With(
		slog.String("op", op),
	)

	if _, ok := db.tasks[task.UUID]; ok {
		log.Error("task alredy in storage")
		return "", fmt.Errorf("task alredy in storage")
	}
	if task.UUID == "" {
		userUUID, _ := uuid.NewUUID()
		task.UUID = userUUID.String()
	}

	db.tasks[task.UUID] = task

	return task.UUID, nil
}

func (db *InMemDB) EditTask(ctx context.Context, task repomodel.Task) (servicemodel.Task, error) {
	const op = "repository.InMemory.EditTask"

	log := db.log.With(
		slog.String("op", op),
	)

	if _, ok := db.tasks[task.UUID]; !ok {
		log.Error("no task in repository")

		return servicemodel.Task{}, fmt.Errorf("no task uuid-%s in repository", task.UUID)
	}

	db.tasks[task.UUID] = task

	return converter.ToTaskFromRepo(task), nil
}
