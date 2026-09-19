package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type employee struct {
	id       int
	name     string
	position string
	salary   int
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

	res, err := db.Exec("INSERT INTO employees (name, position, salary) VALUES ($1, $2, $3)", "Alice", "developer", 120000)
	if err != nil {
		log.Fatal(err)
	}
	resultate, _ := res.RowsAffected()
	fmt.Printf("add %d row(s)\n", resultate)

	res, err = db.Exec("INSERT INTO employees (name, position, salary) VALUES ($1, $2, $3)", "Bob", "designer", 90000)
	if err != nil {
		log.Fatal(err)
	}
	resultate, _ = res.RowsAffected()
	fmt.Printf("add %d row(s)\n", resultate)

	res, err = db.Exec("INSERT INTO employees (name, position, salary) VALUES ($1, $2, $3)", "Charlie", "manager", 150000)
	if err != nil {
		log.Fatal(err)
	}
	resultate, _ = res.RowsAffected()
	fmt.Printf("add %d row(s)\n", resultate)

	res, err = db.Exec("INSERT INTO employees (name, position, salary) VALUES ($1, $2, $3)", "Diana", "developer", 130000)
	if err != nil {
		log.Fatal(err)
	}
	resultate, _ = res.RowsAffected()
	fmt.Printf("add %d row(s)\n", resultate)

	var employeers []employee

	row, err := db.Query("SELECT id, name, position, salary FROM employees")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	for row.Next() {
		var e employee
		err = row.Scan(&e.id, &e.name, &e.position, &e.salary)
		if err != nil {
			log.Fatal(err)
		}
		employeers = append(employeers, e)
	}
	fmt.Println(employeers)

	var pers employee
	err = db.QueryRow("SELECT id, name, position, salary FROM employees WHERE id = $1", 2).Scan(&pers.id, &pers.name, &pers.position, &pers.salary)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found: %d | %s | %s | %d", pers.id, pers.name, pers.position, pers.salary)

	res, err = db.Exec("UPDATE employees SET salary = salary + $1 WHERE position = $2", 10000, "developer")
	if err != nil {
		log.Fatal(err)
	}
	ress, _ := res.RowsAffected()
	fmt.Printf("Update %d row(s)\n", ress)

	res, err = db.Exec("DELETE FROM employees WHERE name = $1", "Charlie")
	if err != nil {
		log.Fatal(err)
	}
	resss, _ := res.RowsAffected()
	fmt.Printf("delete %d row(s)\n", resss)

	var employeers1 []employee

	row, err = db.Query("SELECT id, name, position, salary FROM employees")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	for row.Next() {
		var e employee
		err = row.Scan(&e.id, &e.name, &e.position, &e.salary)
		if err != nil {
			log.Fatal(err)
		}
		employeers1 = append(employeers1, e)
	}
	fmt.Println(employeers1)
}
