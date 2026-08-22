package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

type Cart struct {
	Items []Product
}

func (c Cart) Add(p Product) Cart {
	c.Items = append(c.Items, p)
	return c
}

func (c Cart) Total() float64 {
	sum := 0.0
	for _, Item := range c.Items {
		sum += Item.Price
	}
	return sum
}

func (c Cart) Info() {
	for _, Item := range c.Items {
		fmt.Println(Item.Name)
	}
	fmt.Println("Итого:", c.Total())
}

func main() {
	ProductCart := Cart{}
	ProductCart = ProductCart.Add(Product{"Phone", 20000})
	ProductCart = ProductCart.Add(Product{"TV", 120000})
	ProductCart = ProductCart.Add(Product{"Computer", 80000})
	ProductCart.Info()
}
