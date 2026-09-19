package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type product struct {
	id    int
	name  string
	price int
}

func main() {
	db, err := sql.Open("pgx", "postgres://imac:@localhost:5432/learning?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", "bread", 55)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("\nInsert %d row(s)\n", rowsAffected)

	result2, err := db.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", "milk", 109)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected2, _ := result2.RowsAffected()
	fmt.Printf("\nInsert %d row(s)\n", rowsAffected2)

	result3, err := db.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", "chesee", 250)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected3, _ := result3.RowsAffected()
	fmt.Printf("\nInsert %d row(s)\n", rowsAffected3)

	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []product
	for rows.Next() {
		var p product
		err := rows.Scan(&p.id, &p.name, &p.price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}
	fmt.Println("All Products:")
	for _, product := range products {
		fmt.Printf(" %d | %s | %d", product.id, product.name, product.price)
	}

	_, err = db.Exec("UPDATE products SET price = $1 WHERE name = $2", 60, "bread")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nBread price uptadet to 60")

	_, err = db.Exec("DELETE FROM products WHERE name = $1", "chesee")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nDelete rows chesee")

	rows, err = db.Query("SELECT id, name, price FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products2 []product
	for rows.Next() {
		var p product
		err := rows.Scan(&p.id, &p.name, &p.price)
		if err != nil {
			log.Fatal(err)
		}
		products2 = append(products2, p)
	}
	fmt.Println("All Products:")
	for _, product := range products2 {
		fmt.Printf(" %d | %s | %d", product.id, product.name, product.price)
	}
}
