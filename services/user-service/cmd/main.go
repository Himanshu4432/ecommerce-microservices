// services/user-service/cmd/main.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Himanshu4432/ecommerce-microservices/services/user-service/internal/handlers"
	"github.com/Himanshu4432/ecommerce-microservices/services/user-service/internal/repository"
	"github.com/Himanshu4432/ecommerce-microservices/services/user-service/internal/service"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize Repository
	repo := repository.NewUserRepository()

	// Initialize Service
	userService := service.NewUserService(repo)

	// Initialize Handlers
	userHandlers := handlers.NewUserHandlers(userService)

	// Initialize Router
	router := mux.NewRouter()

	// Define Routes
	router.HandleFunc("/register", userHandlers.Register).Methods("POST")
	router.HandleFunc("/login", userHandlers.Login).Methods("POST")
	router.HandleFunc("/profile/{id}", userHandlers.GetProfile).Methods("GET")

	// Get port from environment or default to 8081
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("User Service is running on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
