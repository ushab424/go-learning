package main

import "fmt"

func ApplyDiscount(price *float64, discount float64) {
	*price = *price * (100 - discount) / 100
}

func main() {
	price := 1000.0
	ApplyDiscount(&price, 10)
	fmt.Println(price)

	ApplyDiscount(&price, 50)
	fmt.Println(price)
}
