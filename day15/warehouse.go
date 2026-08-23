package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type WareHouse struct {
	Products []Product
}

func (w WareHouse) Add(p Product) WareHouse {
	w.Products = append(w.Products, p)
	return w
}

func (w WareHouse) TotalValue() float64 {
	TotalPrice := 0.0
	for _, value := range w.Products {
		TotalPrice += value.Price * float64(value.Quantity)
	}
	return TotalPrice
}

func (w WareHouse) MostExpensive() string {
	var name string
	var price float64
	for _, value := range w.Products {
		if value.Price > price {
			name = value.Name
			price = value.Price
		}
	}
	return name
}

func main() {
	WareHouses := WareHouse{}
	WareHouses = WareHouses.Add(Product{"phone", 11500.00, 3})
	WareHouses = WareHouses.Add(Product{"tv", 111500.30, 1})
	WareHouses = WareHouses.Add(Product{"mouse", 1150.50, 2})
	WareHouses = WareHouses.Add(Product{"table", 121500.00, 1})
	fmt.Println(WareHouses.TotalValue())
	fmt.Println(WareHouses.MostExpensive())
}
