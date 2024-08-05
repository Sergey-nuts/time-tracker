package apimodel

import "time"

type Task struct {
	UUID         string        `json:"uuid,omitempty"`
	Running      bool          `json:"running"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	CreationTime time.Time     `json:"creationtime,omitempty"`
	StartTime    time.Time     `json:"starttime,omitempty"`
	Duration     time.Duration `json:"duration,omitempty"`
}
