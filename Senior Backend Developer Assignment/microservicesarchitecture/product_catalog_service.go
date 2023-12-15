// product_catalog_service.go

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Burada ürünleri getirme işlemleri yapılır.
	fmt.Fprintln(w, "Get Products")
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	// Burada yeni bir ürün oluşturma işlemleri yapılır.
	fmt.Fprintln(w, "Create Product")
}

func main() {
	r := mux.NewRouter()

	// Ürünleri getirme
	r.HandleFunc("/products", GetProductsHandler).Methods("GET")

	// Yeni ürün oluşturma
	r.HandleFunc("/products", CreateProductHandler).Methods("POST")

	// Servisi başlat
	http.Handle("/", r)
	http.ListenAndServe(":8082", nil)
}
