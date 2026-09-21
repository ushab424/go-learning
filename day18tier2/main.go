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

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

var db *sql.DB
var jwtSecret = []byte("our-secret-key-change-this")

func main() {
	var err error
	// connected to database
	db, err = sql.Open("pgx", "postgres://imac:@localhost:5432/learning?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database!")

	// use chi router
	r := chi.NewRouter()

	// request method's
	r.Post("/register", register)
	r.Post("/login", login)
	r.Get("/users/{id}", getUser)
	// use middleware validation
	r.Delete("/users/{id}", authMiddleware(deleteUser))

	fmt.Println("server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// register handler
func register(w http.ResponseWriter, r *http.Request) {
	// create a variable request
	var req RegisterRequest
	// read userInfo from request body, and write to variable
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// translate password to hash
	hash, err := hashPassword(req.Password)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	var userID int
	// create a new row in db with variable data, and scan id to userID from new row
	err = db.QueryRow("INSERT INTO users (name, email, password_hash, age) VALUES ($1, $2, $3, $4) RETURNING id", req.Name, req.Email, hash, req.Age).Scan(&userID)
	if err != nil {
		http.Error(w, "email already exist", http.StatusConflict)
		return
	}

	// create a jwtToken
	token, err := generateToken(userID)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	user := User{ID: userID, Name: req.Name, Email: req.Email, Age: req.Age}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// encode user with AuthResponse func (token+userInfo)
	json.NewEncoder(w).Encode(AuthResponse{Token: token, User: user})
}

// login handler
func login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	// read userInfo from request body, and write to variable
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	var user User
	var hash string
	// read userInfo from db and send data to user and hash variable
	err := db.QueryRow("SELECT id, name, email, age, password_hash FROM users WHERE email = $1", req.Email).Scan(&user.ID, &user.Name, &user.Email, &user.Age, &hash)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	// use verifying func (send hash and request password to function)
	if !verifyPassword(hash, req.Password) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	// genereate a new jwtToken
	token, err := generateToken(user.ID)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	// login done
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthResponse{Token: token, User: user})
}

func getUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var user User
	err = db.QueryRow("SELECT id, name, email, age FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	res, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// middleware - handler in func, we read token, if token == ok => access
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// create variable token with Authorization
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		// validation token with Bearer method
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		// use verify func with request token
		userID, err := verifyToken(token)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// check request id and db id
		idStr := chi.URLParam(r, "id")
		id, _ := strconv.Atoi(idStr)
		if userID != id {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		// if all ok => go next handler
		next(w, r)
	}
}
