package servicemodel

import "time"

type Task struct {
	UUID         string
	Running      bool
	Title        string
	Description  string
	CreationTime time.Time
	StartTime    time.Time
	Duration     time.Duration
}
