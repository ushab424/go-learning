package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Введите три числа:")
	fmt.Scan(&a, &b, &c)
	min := minimum(a, b, c)
	fmt.Println(min)
}
func minimum(a, b, c int) int {
	var min int
	if b < a {
		min = b
	} else {
		min = a
	}
	if c < min {
		min = c
	}
	return min
}
