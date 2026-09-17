package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// create a User struct
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// create a UserList
var users = []User{
	{ID: 1, Name: "Ivan", Email: "ivan@mail.com"},
	{ID: 2, Name: "Olga", Email: "olga@mail.com"},
}

// Next id == 3
var nextID = 3

// function show all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// func show one user from userList with id
func getUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, u := range users {
		if u.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
}

// function create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u.ID = nextID
	nextID++
	users = append(users, u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

// function delete users with his id
func deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	r := chi.NewRouter()

	r.Get("/users", getUsers)
	r.Get("/user/{id}", getUser)
	r.Post("/users", createUser)
	r.Delete("/users/{id}", deleteUser)

	fmt.Println("server startes on :8080")
	http.ListenAndServe(":8080", r)
}
