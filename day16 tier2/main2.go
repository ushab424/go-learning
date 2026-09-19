package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	db, err := sql.Open("pgx", "postgres://imac:@localhost:5432/learning?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// SELECT all — сканируем в структуру
	rows, err := db.Query("SELECT id, name, email, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	fmt.Println("All users:")
	for _, u := range users {
		fmt.Printf("  %d | %s | %s | %d\n", u.ID, u.Name, u.Email, u.Age)
	}
}
