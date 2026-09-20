package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	strID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	err = db.QueryRow("SELECT id, name, price FROM products WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "product not found"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	err = db.QueryRow("INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id", product.Name, product.Price).Scan(&product.ID)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func deleteProducts(w http.ResponseWriter, r *http.Request) {
	strID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	res, err := db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "product not found"})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	db, err := sql.Open("pgx", "postgres://imac:@localhost:5432/learning?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Get("/products", getProducts)
	r.Get("/products/{id}", getProduct)
	r.Post("/products", createProduct)
	r.Delete("/products/{id}", deleteProducts)

	fmt.Println("server started on :8080")
	http.ListenAndServe(":8080", r)
}
