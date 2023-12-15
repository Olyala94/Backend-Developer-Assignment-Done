package main

import (
	"fmt"
	"net/http"
)

func handleTasks(w http.ResponseWriter, r *http.Request) {
	// Task listesini döndürme veya diğer işlemleri gerçekleştirme
	fmt.Fprintf(w, "Task Listesi Burada Olacak!")
}

func main() {
	// HTTP server'ı başlatma
	http.HandleFunc("/tasks", handleTasks)
	http.ListenAndServe(":8080", nil)
}
