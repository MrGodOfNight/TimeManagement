package main

import (
	"TimeManagement/src/server/model"
	"TimeManagement/src/server/route"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// Download .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	// Connect to database
	model.ConnectDB()

	// Create HTTP server
	mux := http.NewServeMux()

	// Register routes
	route.Routes(mux)

	// Start server
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
