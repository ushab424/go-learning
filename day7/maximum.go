package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Введите два числа:")
	fmt.Scan(&a, &b)
	maximum := max(a, b)
	fmt.Println(maximum)
}
func max(a, b int) int {
	maximum := 0
	if a > b {
		maximum = a
	} else {
		maximum = b
	}
	return maximum
}
