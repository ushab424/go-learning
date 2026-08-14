package main

import "fmt"

func main() {
	var a float64
	var b string
	fmt.Println("введите температуру и единицу измерения:")
	fmt.Scan(&a, &b)
	if b == "C" {
		fmt.Printf("%.2f F", a*1.8+32)
	} else if b == "F" {
		fmt.Printf("%.2f C", (a-32)/1.8)
	} else {
		fmt.Printf("Введена неправильная единица измерения\n")
	}
}
