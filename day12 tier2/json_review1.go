package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Pages   int    `json:"pages,omitempty"`
	InStock bool   `json:"in_stock"`
}

func main() {
	Books := []Book{
		{Title: "BlackHatGo", Author: "Ivan Dorn", Pages: 130, InStock: true},
		{Title: "Kimono", Author: "Dustin Bomber", Pages: 0, InStock: true},
		{Title: "BibiThisMouth", Author: "Clow Sniper", Pages: 3476, InStock: false},
	}
	// Marshall all slice
	data, err := json.MarshalIndent(Books, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(data))
	// Unmarshall
	var parsed []Book
	err = json.Unmarshal(data, &parsed)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	for _, book := range parsed {
		fmt.Printf("Title: %s, Author: %s, Pages: %d, InStock: %t\n", book.Title, book.Author, book.Pages, book.InStock)
	}
}
