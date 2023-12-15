package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Merhaba, %s!", r.URL.Path[1:])
}

func main2() {
	http.HandleFunc("/tasks", handleTasks)
	http.HandleFunc("/tasks/:id", handleTaskByID)

	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	// Handle GET and POST requests for tasks
}

func handleTaskByID(w http.ResponseWriter, r *http.Request) {
	// Handle GET, PUT, and DELETE requests for a specific task
}
