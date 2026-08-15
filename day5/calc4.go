package main

import "fmt"

func main() {
	var a, b float64
	var c string
	fmt.Println("Введите пример:")
	fmt.Scan(&a, &c, &b)
	switch c {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя!")
		} else {
			fmt.Println(a / b)
		}
	default:
		fmt.Println("Неизвестная операция")
	}
}
