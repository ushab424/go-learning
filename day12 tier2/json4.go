package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User1 struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// Запись в файл
	users := []User1{
		{Name: "Ivan", Age: 15, Email: "ivan@mail.com"},
		{Name: "Olga", Age: 25, Email: "olga@mail.com"},
	}

	file, err := os.Create("users.json")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(users)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("file written")
	// Чтение из файла
	file2, err := os.Open("users.json")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer file2.Close()

	var readUsers []User1
	decoder := json.NewDecoder(file2)
	err = decoder.Decode(&readUsers)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, u := range readUsers {
		fmt.Printf("Name: %s, Age: %d, Email: %s\n", u.Name, u.Age, u.Email)
	}

}
