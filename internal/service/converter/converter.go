package converter

import (
	"time-tracker/internal/repository/repomodel"
	"time-tracker/internal/service/servicemodel"
)

func TaskToRepoFromService(task servicemodel.Task) repomodel.Task {
	return repomodel.Task{
		UUID:         task.UUID,
		Title:        task.Title,
		Description:  task.Description,
		CreationTime: task.CreationTime,
	}
}
