package converter

import (
	"time-tracker/internal/repository/repomodel"
	"time-tracker/internal/service/servicemodel"
)

func ToTaskFromRepo(task repomodel.Task) servicemodel.Task {
	return servicemodel.Task{
		UUID:         task.UUID,
		Running:      task.Running,
		Title:        task.Title,
		Description:  task.Description,
		CreationTime: task.CreationTime,
		StartTime:    task.StartTime,
		Duration:     task.Duration,
	}
}
