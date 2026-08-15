package main

import "fmt"

func main() {
	var a, b float64
	fmt.Print("введите два числа:")
	fmt.Scan(&a, &b)
	fmt.Println("сумма=", a+b)
	fmt.Println("разность=", a-b)
	fmt.Println("произведение=", a*b)
	fmt.Println("частное=", a/b)
}
