package converter

import (
	"time-tracker/internal/repository/repomodel"
	"time-tracker/internal/server/apimodel"
	"time-tracker/internal/service/servicemodel"
)

// TaskToRepo convert task from service model to repository model
func TaskToRepo(task servicemodel.Task) repomodel.Task {
	return repomodel.Task{
		UUID:         task.UUID,
		Running:      task.Running,
		Title:        task.Title,
		Description:  task.Description,
		CreationTime: task.CreationTime,
		StartTime:    task.StartTime,
		Duration:     task.Duration,
	}
}

// TaskToApi converts task from service model to api model
func TaskToApi(task servicemodel.Task) apimodel.Task {
	return apimodel.Task{
		UUID:         task.UUID,
		Running:      task.Running,
		Title:        task.Title,
		Description:  task.Description,
		CreationTime: task.CreationTime,
		StartTime:    task.StartTime,
		Duration:     task.Duration,
	}
}
