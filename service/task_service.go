package service

import (
	"fmt"
	"math/rand"
	"sync"
	"task-manager/models"
	"time"

	"github.com/google/uuid"
)

type TaskService struct {
	task map[string]*models.Task
	mutex sync.RWMutex
	workers chan struct{}
}

func NewTaskService() *TaskService {
	return &TaskService{
		task: make(map[string]*models.Task),
		workers: make(chan struct{}, 5),
	}
}

func (s *TaskService) CreateTask() *models.Task {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	id := uuid.New().String()
	task := &models.Task{
		ID: id,
		Status: models.StatusPending,
		TimeAtCreated: time.Now(),
		Result: "",
		Error: "",
	}
	s.task[id] = task

	go s.executeTask(task)

	return task
}

func (s *TaskService) GetTask(id string) (*models.Task, error){
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	task, exists := s.task[id] 
	if !exists {
		return nil, fmt.Errorf("task not found")
	}
	copyTask := *task
	if task.TimeAtStarted != nil {
		if copyTask.TimeAtStarted != nil {
			var duration time.Duration
			if copyTask.TimeAtCompleted != nil {
				duration = copyTask.TimeAtCompleted.Sub(*copyTask.TimeAtStarted)
			} else {
				duration = time.Now().Sub(*copyTask.TimeAtStarted)
			}
			copyTask.InProcess = duration.String()
		}
	}
	return &copyTask, nil
}

func (s *TaskService) DeleteTask(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.task[id]; !exists {
		return fmt.Errorf("task not found")
	}

	delete(s.task, id)
	return nil

}

func (s *TaskService) GetAllTask() []*models.Task {

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	tasks := make([]*models.Task, 0, len(s.task))
	
	for _, task := range s.task{
		copyTask := *task
		if copyTask.TimeAtStarted != nil {
			var duration time.Duration
			if copyTask.TimeAtCompleted != nil {
				duration = copyTask.TimeAtCompleted.Sub(*copyTask.TimeAtStarted)
			} else {
				duration = time.Now().Sub(*copyTask.TimeAtStarted)
			}
			copyTask.InProcess = duration.String()
		}
		tasks = append(tasks, &copyTask)
	}
	return tasks
}

func (s *TaskService) executeTask(task *models.Task) {
	s.workers <- struct{}{}
	defer func(){ <- s.workers}()

	s.mutex.Lock()
	task.Status = models.StatusRunning
	now := time.Now()
	task.TimeAtStarted = &now
	s.mutex.Unlock()

	err := s.simulatorIOWork()

	s.mutex.Lock()
	completedTime := time.Now()
	task.TimeAtCompleted = &completedTime

	if err != nil {
		task.Status = models.StatusFailed
		task.Error = err.Error()
	} else {
		task.Status = models.StatusComplete
		task.Result = "Task completed successfully"
	}
	s.mutex.Unlock()
}

func (s* TaskService) simulatorIOWork() error {
	duration := time.Duration(3+rand.Intn(3)) * time.Minute
	time.Sleep(duration)

	if rand.Intn(100) < 20 {
		return fmt.Errorf("simulation failed: network timeout")
	}

	return nil
}