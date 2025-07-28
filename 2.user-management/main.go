package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/neel07sanghvi/crud-api/handlers"
	"github.com/neel07sanghvi/crud-api/storage"
)

func main() {
	userStorage := storage.New()

	userHandler := handlers.New(userStorage)

	http.HandleFunc("/users", userHandler.HandleUsers)
	http.HandleFunc("/users/", userHandler.HandleUsers)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status": "healthy"}`)
	})

	port := "8080"

	fmt.Printf("Server starting on port %s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET    /health        - Health check")
	fmt.Println("  GET    /users         - Get all users")
	fmt.Println("  GET    /users/1       - Get user by ID")
	fmt.Println("  POST   /users         - Create new user")
	fmt.Println("  PUT    /users/1       - Update user")
	fmt.Println("  DELETE /users/1       - Delete user")
	fmt.Println()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
