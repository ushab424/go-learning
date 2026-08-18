package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Введите сумму и степень:")
	fmt.Scan(&a, &b)
	full := power(a, b)
	fmt.Println(full)
}
func power(a, b int) int {
	if b == 0 {
		return 1
	}
	return a * power(a, b-1)
}
