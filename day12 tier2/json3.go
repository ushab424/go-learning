package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
}

type Employee struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

func main() {
	Employee := Employee{Name: "Ivan", Age: 15, Address: Address{City: "Moscow", Street: "Mishurina street"}}
	EmpData1, err := json.Marshal(Employee)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(EmpData1))
}
