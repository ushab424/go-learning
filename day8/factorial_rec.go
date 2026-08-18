package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите число:")
	fmt.Scan(&a)
	factor := fac(a)
	fmt.Println(factor)
}
func fac(a int) int {
	if a == 1 {
		return a
	}
	return a * fac(a-1)
}
