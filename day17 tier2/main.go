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

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

var db *sql.DB

// var db *sql.DB — глобальная переменная для подключения к базе. Объявляем глобально, чтобы все хендлеры имели доступ к базе.

// request handler all users from PostgreSQL to Browser
func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, email, age FROM users")
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}
	// this rows.err check error from iteration
	if err = rows.Err(); err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// request handler one row from PostgreSQL with id
func getUser(w http.ResponseWriter, r *http.Request) {
	var person User
	strID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "request error", http.StatusBadRequest)
		return
	}
	err = db.QueryRow("SELECT id, name, email, age FROM users WHERE id = $1", id).Scan(&person.ID, &person.Name, &person.Email, &person.Age)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

// this hadler create a new user in to PostgreSQL
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := db.QueryRow("INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id", user.Name, user.Email, user.Age).Scan(&user.ID)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	idUser := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idUser)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	res, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

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

	r.Get("/users", getUsers)
	r.Get("/users/{id}", getUser)
	r.Post("/users", createUser)
	r.Delete("/users/{id}", deleteUser)

	fmt.Println("server strated on :8080")
	http.ListenAndServe(":8080", r)
}
