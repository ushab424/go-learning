package main

import "fmt"

func main() {
	var a float64
	var b float64
	var c string
	fmt.Println("введите пример:")
	fmt.Scan(&a, &c, &b)
	if c == "+" {
		fmt.Println(a + b)
	} else if c == "-" {
		fmt.Println(a - b)
	} else if c == "*" {
		fmt.Println(a * b)
	} else if c == "/" {
		if b == 0 {
			fmt.Println("на ноль делить нельзя")
		} else {
			fmt.Println(a / b)
		}
	}
}
