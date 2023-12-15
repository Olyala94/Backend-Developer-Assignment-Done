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

	// Güvenlik kontrolleri
	if err := performSecurityChecks(r); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// İşlemler
	fmt.Fprintf(w, "Merhaba, %s!", r.URL.Path)
}

func performSecurityChecks(r *http.Request) error {
	// Implemente edilecek: Güvenlik kontrollerini gerçekleştirme
	return nil
}
