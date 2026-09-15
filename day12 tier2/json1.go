package main

import (
	"encoding/json"
	"fmt"
)

// JSON - JavaScript Object Notation

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// JSON-теги — то что ты написал в обратных кавычках

// Ключ — всегда строка в кавычках. Значение — строка, число, bool, массив, объект или null.

func main() {
	user := User{Name: "Ivan", Age: 15, Email: "Ervisto967@gmail.com"}
	data, err := json.Marshal(user)
	// // data = []byte(`{"name":"Ivan","age":15,"email":"test@mail.com"}`)

	// json.Marshal — берёт Go-структуру и превращает в JSON (в виде []byte).
	// Слово "маршалинг" = сериализация = "упаковать данные в формат для передачи".
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(data))

	jsonStr := `{"name":"Olga","age":25,"email":"olga@mail.com"}`
	var user2 User
	err = json.Unmarshal([]byte(jsonStr), &user2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(user2.Name, user2.Age, user2.Email)
}
