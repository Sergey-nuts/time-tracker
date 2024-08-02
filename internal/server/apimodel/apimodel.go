package apimodel

import "time"

type TaskInfo struct {
	UUID string `json:"uuid,omitempty"`
	//Running      bool	`json:"running"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreationTime time.Time `json:"creationtime,omitempty"`
	//StartTime    time.Time
	//Duration     time.Duration
}
