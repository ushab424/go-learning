package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	// это драйвер PostgreSQL. Нижнее подчёркивание _ означает, что мы не вызываем его напрямую,
	// но при импорте он регистрирует себя внутри database/sql. После этого sql.Open("pgx", ...) знает, как говорить с PostgreSQL.
)

func main() {
	connStr := "postgres://imac:@localhost:5432/learning?sslmode=disable"
	db, err := sql.Open("pgx", connStr)
	// открывает пул соединений к базе. Первый аргумент — имя драйвера,
	// второй — строка подключения: postgres://имя_пользователя:пароль@хост:порт/имя_базы.
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// закроет соединение при завершении программы.

	err = db.Ping()
	// проверяет, что соединение реально работает. sql.Open само по себе не подключается, а только готовит пул.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database!")

	// SELECT - import all rows from table users
	rows, err := db.Query("SELECT id, name, email, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// rows — это набор строк из базы, как курсор
	// defer rows.Close() — обязательно закрываем после использования

	fmt.Println("\nAll users:")

	for rows.Next() {
		// rows.Next() — переходит к следующей строке, возвращает false когда строки кончились
		var id int
		var name string
		var email string
		var age int

		err := rows.Scan(&id, &name, &email, &age)
		// rows.Scan — копирует значения текущей строки в переменные
		// порядок переменных должен совпадать с порядком столбцов в SELECT
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d | Name: %s | Email: %s | Age: %d\n", id, name, email, age)
	}

	fmt.Println("\nSearch user with id = 2:")

	var u_id int
	var u_name string
	var u_email string
	var u_age int

	err = db.QueryRow("SELECT id, name, email, age FROM users WHERE id = $1", 2).Scan(&u_id, &u_name, &u_email, &u_age)
	// db.QueryRow — то же, что db.Query, но возвращает ровно одну строку. Не нужен цикл for rows.Next(), сразу вызываешь .Scan().

	// $1 — плейсхолдер. Вместо того чтобы вставлять значение прямо в строку запроса (WHERE id = 2),
	// мы пишем $1 и передаём значение вторым аргументом. PostgreSQL использует $1, $2, $3 для параметров.
	// Это защищает от SQL-инъекций — атак, когда злоумышленник подставляет свой SQL-код через пользовательский ввод.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %d | Name: %s | Email: %s | Age: %d\n", u_id, u_name, u_email, u_age)

	result, err := db.Exec("INSERT INTO users (name, email, age) VALUES ($1, $2, $3)", "Sergey", "sergey@gmail.com", 28)
	// db.Exec(sql, args...) — выполняет запрос, который не возвращает строки. Используется для INSERT, UPDATE, DELETE. Возвращает result и error.

	// $1, $2, $3 — плейсхолдеры, значения передаются по порядку после SQL-строки: "Sergey" подставится в $1, "sergey@mail.com" в $2, 28 в $3.
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()
	// result.RowsAffected() — сколько строк было затронуто. Для INSERT это 1 (вставили одну строку), для DELETE — сколько удалили.
	fmt.Printf("\nInserted %d row(s)\n", rowsAffected)

	result, err = db.Exec("DELETE FROM users WHERE name = $1", "Sergey")
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ = result.RowsAffected()
	fmt.Printf("Deletes %d row(s)\n", rowsAffected)
}
