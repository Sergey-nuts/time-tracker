package converter

import (
	"time-tracker/internal/server/apimodel"
	"time-tracker/internal/service/servicemodel"
)

// TaskToService converts task from api model to service model
func TaskToService(task apimodel.Task) servicemodel.Task {
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
