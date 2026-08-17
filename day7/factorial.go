package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите число:")
	fmt.Scan(&a)
	fac := factorial(a)
	fmt.Println(fac)
}
func factorial(x int) int {
	factor := 1
	for i := 1; i <= x; i++ {
		factor = factor * i
	}
	return factor
}
