package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "running", "version": "1.0"})
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, err1 := strconv.Atoi(aStr)
	b, err2 := strconv.Atoi(bStr)

	if err1 != nil || err2 != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid parametrs"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"a": a, "b": b, "sum": a + b})
}

func main() {
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/calculate", calculateHandler)

	fmt.Println("server start on :8080")
	http.ListenAndServe(":8080", nil)
}
