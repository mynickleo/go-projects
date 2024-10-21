package main

import (
	"log"
	"net/http"
	"todo-backend/internal/tasks"

	"github.com/gorilla/mux"
)

func main() {
	repo, err := tasks.NewRepository()
	if err != nil {
		log.Fatal(err)
	}

	handler := &tasks.Handler{Repo: repo}
	r := mux.NewRouter()

	r.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", handler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", handler.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
