package converter

import (
	"time-tracker/internal/server/apimodel"
	"time-tracker/internal/service/servicemodel"
)

func ToTaskFromApi(in apimodel.Task) servicemodel.Task {
	return servicemodel.Task{
		UUID:         in.UUID,
		Running:      in.Running,
		Title:        in.Title,
		Description:  in.Description,
		CreationTime: in.CreationTime,
		StartTime:    in.StartTime,
		Duration:     in.Duration,
	}
}
