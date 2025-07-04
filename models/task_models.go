package models

import "time"

type TaskStatus = string

const (
	StatusPending  TaskStatus = "pending"
	StatusRunning  TaskStatus = "running"
	StatusComplete TaskStatus = "completed"
	StatusFailed   TaskStatus = "failed"
)

type Task struct {
	ID              string     `json:"id"`
	Status          TaskStatus `json:"status"`
	TimeAtCreated   time.Time  `json:"time_at_created"`
	TimeAtStarted   *time.Time `json:"time_at_started"`
	TimeAtCompleted *time.Time `json:"time_at_completed"`
	Result          string     `json:"result"`
	Error           string     `json:"error"`
	InProcess       string     `json:"in_process"`
}
