package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"task-manager/service"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler (taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}


func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {

	task := h.taskService.CreateTask()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if path == ""{
		http.Error(w, "Bad ID", http.StatusBadRequest)
		return
	}

	task, err := h.taskService.GetTask(path)
	if err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)

}

func (h *TaskHandler) GetAllTask(w http.ResponseWriter, r *http.Request) {

	tasks := h.taskService.GetAllTask()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if path == "" {
		http.Error(w, "Bad ID", http.StatusBadRequest)
		return
	}

	err := h.taskService.DeleteTask(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}