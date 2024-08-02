package repomodel

import "time"

type Task struct {
	UUID         string
	Running      bool
	User_id      string
	Title        string
	Description  string
	CreationTime time.Time
	StartTime    time.Time
	Duration     time.Duration
}
