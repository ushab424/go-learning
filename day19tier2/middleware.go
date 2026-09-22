package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// loggingMiddleware - logging all HTTP requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// call to next handler
		next.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("%s %s %s %dms\n", r.Method, r.RequestURI, r.RemoteAddr, duration.Milliseconds())
	})
}

// recoverMiddleware - catch all panic and send statusCode 500
func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Panic %v\n", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "missed token", http.StatusUnauthorized)
			return
		}

		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		UserID, err := verifyToken(token)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		idStr := strings.TrimPrefix(r.URL.Path, "/users/")
		id, _ := strconv.Atoi(idStr)
		if UserID != id {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}
