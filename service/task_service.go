package service

import (
	"sync"
	"task-manager/models"
)

type TaskService struct {
	task map[string]*models.Task
	mutex sync.Mutex
	workers chan struct{}
}

func NewTaskService() *TaskService {
	return &TaskService{
		task: make(map[string]*models.Task),
		workers: make(chan struct{}, 5),
	}
}