package converter

import (
	"time-tracker/internal/server/apimodel"
	"time-tracker/internal/service/servicemodel"
)

func ToTaskFromApi(in apimodel.TaskInfo) servicemodel.Task {
	return servicemodel.Task{
		UUID:         in.UUID,
		Title:        in.Title,
		Description:  in.Description,
		CreationTime: in.CreationTime,
	}
}
