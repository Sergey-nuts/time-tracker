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

func (db *InMemDB) Tasks(_ context.Context, _ int) ([]servicemodel.Task, error) {
	const op = "repository.InMemory.Tasks"

	log := db.log.With(
		slog.String("op", op),
	)
	_ = log

	t := make([]servicemodel.Task, 0, len(db.tasks))
	for _, v := range db.tasks {
		t = append(t, converter.TaskToService(v))
	}

	return t, nil
}

func (db *InMemDB) AddTask(_ context.Context, task repomodel.Task) (servicemodel.Task, error) {
	const op = "repository.InMemory.AddTask"

	log := db.log.With(
		slog.String("op", op),
	)

	if _, ok := db.tasks[task.UUID]; ok {
		log.Error("task alredy in storage")
		return servicemodel.Task{}, fmt.Errorf("task alredy in storage")
	}
	if task.UUID == "" {
		UUID, _ := uuid.NewUUID()
		task.UUID = UUID.String()
	}

	db.tasks[task.UUID] = task

	return converter.TaskToService(task), nil
}

func (db *InMemDB) EditTask(ctx context.Context, task repomodel.Task) (servicemodel.Task, error) {
	const op = "repository.InMemory.EditTask"

	log := db.log.With(
		slog.String("op", op),
	)

	if _, ok := db.tasks[task.UUID]; !ok {
		log.Error("no task in repository")

		return servicemodel.Task{}, repository.ErrorTaskNotFound
	}

	db.tasks[task.UUID] = task

	return converter.TaskToService(task), nil
}

func (db *InMemDB) DeleteTask(ctx context.Context, task repomodel.Task) (servicemodel.Task, error) {
	const op = "repository.InMemory.Delete"

	log := db.log.With(
		slog.String("op", op),
	)

	if _, ok := db.tasks[task.UUID]; !ok {
		log.Error("no task in repository")

		return servicemodel.Task{}, repository.ErrorTaskNotFound
	}

	deleted := db.tasks[task.UUID]
	delete(db.tasks, task.UUID)

	return converter.TaskToService(deleted), nil
}

func (db *InMemDB) Get(ctx context.Context, uuid string) (servicemodel.Task, error) {
	const op = "repository.InMemory.Get"

	log := db.log.With(
		slog.String("op", op),
	)

	t, ok := db.tasks[uuid]
	if !ok {
		log.Error("no task in repository")

		return servicemodel.Task{}, repository.ErrorTaskNotFound
	}

	return converter.TaskToService(t), nil

}
