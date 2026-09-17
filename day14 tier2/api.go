package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Movie struct {
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Rating float64 `json:"rating"`
}

var movies = []Movie{
	{Title: "Inception", Year: 2010, Rating: 8.8},
	{Title: "Matrix", Year: 1999, Rating: 8.7},
	{Title: "Interstellar", Year: 2014, Rating: 8.6},
}

func moviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func movieHandler(w http.ResponseWriter, r *http.Request) {
	film := r.URL.Query().Get("title")
	for _, m := range movies {
		if m.Title == film {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(m)
			return
		}

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "movie not found"})
}

func main() {
	http.HandleFunc("/movies", moviesHandler)
	http.HandleFunc("/movie", movieHandler)

	fmt.Println("server started on :8080")
	http.ListenAndServe(":8080", nil)
}
