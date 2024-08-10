/*
	MIT License

	Copyright (c) 2024 Ushakov Igor

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.

*/

package main

import (
	"TimeManagement/src/server/model"
	"TimeManagement/src/server/route"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Create logger
	logger, err := model.NewLogger()
	if err != nil {
		log.Fatalf("Error creating logger: %v\n", err)
	}

	// Download .env file
	err = godotenv.Load()
	if err != nil {
		logger.Error(true, "Error loading .env file: \n", err)
		os.Exit(1)
	}

	// Connect to database
	model.ConnectDB()

	// Create HTTP server
	mux := http.NewServeMux()

	// Register routes
	route.Routes(mux)

	// Create super admin user
	//route.Admin("admin", "admin")

	// Start server
	logger.Info(false, "Starting server at http://"+os.Getenv("SERVER_URL"))
	if err := http.ListenAndServe(os.Getenv("SERVER_URL"), mux); err != nil {
		logger.Error(true, "Error starting server: \n", err)
		os.Exit(1)
	}
}
