package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Order struct {
	ID       int      `json:"id"`
	Total    float64  `json:"total"`
	Customer Customer `json:"customer"`
}

// create json-tags

type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	orders := []Order{
		// create 2 orders
		{ID: 1, Total: 1000, Customer: Customer{Name: "Ivan", Email: "IvanKoehler@gmail.com"}},
		{ID: 2, Total: 2000, Customer: Customer{Name: "Boris", Email: "BorisJohnson@mail.com"}},
	}
	fileOrders, err := os.Create("orders.json")
	// create a json file
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer fileOrders.Close()

	encoder := json.NewEncoder(fileOrders)
	// create a new json encoder
	encoder.SetIndent("", " ")
	// beautiful intedent
	err = encoder.Encode(orders)
	// error encoder
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("File written ok")

	// decoding jsonFile in struct
	fileOrdersPrint, err := os.Open("orders.json")
	// open file
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer fileOrdersPrint.Close()

	var readOrder []Order
	// create a new slice in struct
	decoder := json.NewDecoder(fileOrdersPrint)
	// create a new decoder
	err = decoder.Decode(&readOrder)
	// func decoding
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	for _, Ord := range readOrder {
		// print decoding
		fmt.Println(Ord)
	}
}
