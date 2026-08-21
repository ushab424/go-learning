package main

import "fmt"

type Address struct {
	City   string
	Street string
}

type Employee struct {
	name    string
	age     int
	Address Address
}

func main() {
	Employee1 := Employee{
		name: "Ivan",
		age:  25,
		Address: Address{
			City:   "Moscow",
			Street: "Rostokinskaya",
		},
	}
	fmt.Println(Employee1.name, Employee1.Address.City)
}
