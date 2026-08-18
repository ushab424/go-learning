package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите число:")
	fmt.Scan(&a)
	max := sum(a)
	fmt.Println(max)
}
func sum(a int) int {
	if a == 1 {
		return a
	}
	return a + sum(a-1)
}
