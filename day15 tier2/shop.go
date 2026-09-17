package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []Product{
	{ID: 1, Name: "Bread", Price: 55.30},
	{ID: 2, Name: "Milk", Price: 109.50},
	{ID: 3, Name: "Cabage", Price: 80.00},
}
var NextID = 4

// show all productList
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// show one of products with his id
func getProduct(w http.ResponseWriter, r *http.Request) {
	idProd := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idProd)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, p := range products {
		if id == p.ID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "product not found"})
}

// create a new product in productList
func newProduct(w http.ResponseWriter, r *http.Request) {
	var prod Product
	if err := json.NewDecoder(r.Body).Decode(&prod); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	prod.ID = NextID
	NextID++
	products = append(products, prod)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(prod)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	r := chi.NewRouter()

	r.Get("/products", getProducts)
	r.Get("/products/{id}", getProduct)
	r.Post("/product", newProduct)
	r.Delete("/products/{id}", deleteProduct)

	fmt.Println("server started on :8080")
	http.ListenAndServe(":8080", r)
}
