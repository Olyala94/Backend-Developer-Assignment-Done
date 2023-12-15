package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex // sync.Mutex tanımlanması

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Günlüğe kayıt ekleme
	log.Printf("İstek alındı: %s", r.URL.Path)

	// İşlemler
	fmt.Fprintf(w, "Merhaba, %s!", r.URL.Path)
}

func main() {
	// HTTP server'ı başlatma
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
