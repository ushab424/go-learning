package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите число:")
	fmt.Scan(&a)
	fibonachi := fib(a)
	fmt.Println(fibonachi)
}
func fib(a int) int {
	if a == 0 {
		return 0
	}
	if a == 1 {
		return 1
	}
	return fib(a-1) + fib(a-2)
}
