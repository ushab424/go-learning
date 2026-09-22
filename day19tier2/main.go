package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"day19tier2/handler"
	"day19tier2/repository"
	"day19tier2/service"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/cors"
)

var db *sql.DB
var jwtSecret = []byte("your-secret-key-change-this")

func main() {
	var err error
	db, err = sql.Open("pgx", "postgres://imac:@localhost:5432/learning?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// init all level
	repo := repository.NewUserRepository(db) // work with database
	svc := service.NewUserService(repo)      // buisness logic
	h := handler.NewUserHandler(svc)         // handler

	r := chi.NewRouter()

	// Middleware
	r.Use(loggingMiddleware)      // global middleware logging
	r.Use(cors.Default().Handler) // CORS middleware
	r.Use(recoverMiddleware)      // recover middleware

	r.Post("/register", h.Register)
	r.Post("/login", h.Login)
	r.Get("/users/{id}", h.GetUser)
	r.Delete("/users/{id}", authMiddleware(h.DeleteUser))

	// Graceful shutdown
	// after signal Ctrl+C server correct stopped
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		fmt.Println("Server started on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// wait signal "off"
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\nShutting down server...")

	// get server 5 seconds to exit all active request
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server gracefully stopped")
}
