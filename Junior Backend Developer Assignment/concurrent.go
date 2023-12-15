package main

import (
	"fmt"
	"net/http"
	"sync"
)

var mu sync.Mutex

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// İşlemler
	fmt.Fprintf(w, "Merhaba, %s!", r.URL.Path[1:])
}
