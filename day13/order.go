package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

type OrderProduct struct {
	ID       int
	Product  Product
	Quantity int
}

func main() {
	UserOrder := []OrderProduct{
		{ID: 2047569, Product: Product{Name: "Phone", Price: 2500}, Quantity: 2},
		{ID: 2979375, Product: Product{Name: "Laptop", Price: 12500}, Quantity: 1},
		{ID: 4738758, Product: Product{Name: "iPad", Price: 5000}, Quantity: 3},
	}
	for _, result := range UserOrder {
		fmt.Printf("%d: %s * %d = %d\n", result.ID, result.Product.Name, result.Quantity, (result.Product.Price * result.Quantity))
	}
}
