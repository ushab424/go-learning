package main

import (
	"database/sql"
	"day20tier2/handler"
	"day20tier2/repository"
	"day20tier2/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("pgx", "postgres://imac:@localhost:5432/learning?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	repo := repository.NewProductRepository(db)
	svc := service.NewProductService(repo)
	h := handler.NewProductHandler(svc)

	r.Get("/products/{id}", h.GetProduct)
	r.Post("/products", h.AddProduct)
	r.Delete("/products/{id}", h.DeleteProduct)

	log.Println("server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
