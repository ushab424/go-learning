package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func statusingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "runing", "version": "1.0"})
}

func calculatingHandler(w http.ResponseWriter, r *http.Request) {
	strA := r.URL.Query().Get("a")
	strB := r.URL.Query().Get("b")

	a, erra := strconv.Atoi(strA)
	b, errb := strconv.Atoi(strB)

	if erra != nil || errb != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"run": a, "ver": b, "sum": a + b})
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/calculating", calculatingHandler)

	fmt.Println("server started on :8080 port")
	http.ListenAndServe(":8080", nil)
}
