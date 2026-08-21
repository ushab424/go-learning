package main

import "fmt"

type product struct {
	Name   string
	Price  int
	InStok bool
}

func main() {
	catalog := []product{
		{Name: "Phone", Price: 2000, InStok: true},
		{Name: "Laptop", Price: 5000, InStok: false},
		{Name: "Headphones", Price: 500, InStok: true},
	}
	for _, item := range catalog {
		if item.InStok == true {
			fmt.Println(item.Name, item.Price)
		}
	}
}
