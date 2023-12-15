package main

import (
	"fmt"
	"net/http"
	"sync" // sync paketini import etmeyi unutmayın
)

var mu sync.Mutex // Mutex tanımlanması

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// İstek doğrulama
	if err := validateRequest(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// İşlemler
	fmt.Fprintf(w, "Merhaba, %s!", r.URL.Path[1:])
}

func validateRequest(r *http.Request) error {
	// Implemente edilecek: Gelen verileri doğrulama
	return nil
}
