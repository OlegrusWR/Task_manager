package main

import (
	"log"
	"net/http"
	"task-manager/handlers"
	"task-manager/service"
)

func main () {
	taskService := service.NewTaskService()
	taskHandler := handlers.NewTaskHandler(taskService)

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodPost:
			taskHandler.CreateTask(w, r)
		case http.MethodGet:
			taskHandler.GetAllTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}	
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request){
		switch r.Method{
		case http.MethodGet:
			taskHandler.GetTask(w, r)
		case http.MethodDelete:
			taskHandler.DeleteTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server starting on: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}