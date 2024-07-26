package main

import (
	"TimeManagement/src/server/model"
	"TimeManagement/src/server/route"
	"fmt"
	"net/http"
)

func main() {
	model.ConnectDB()
	mux := http.NewServeMux()
	route.Routes(mux)
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
