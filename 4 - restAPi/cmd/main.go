package main

import (
	"encoding/json"
	"log"
	"net/http"
	"restApi/internal/handlers"
	"restApi/internal/repository"
	"restApi/internal/service"
	"restApi/pkg"

	"github.com/gorilla/mux"
)

func main() {
	db := pkg.InitDB()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/", HomeHandler).Methods("GET")

	log.Printf("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"message": "Hello, World!",
	}

	json.NewEncoder(w).Encode(response)
}
