package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Введите два числа:")
	fmt.Scan(&a, &b)
	result := add(a, b)
	fmt.Println(result)
}
func add(a, b int) int {
	var sum int
	sum = a + b
	return sum
}
