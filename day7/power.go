package main

import "fmt"

func main() {
	var base, exp int
	fmt.Println("Введите число и степень:")
	fmt.Scan(&base, &exp)
	step := power(base, exp)
	fmt.Println(step)
}
func power(base, exp int) int {
	stepen := 1
	for i := 1; i <= exp; i++ {
		stepen *= base
	}
	return stepen
}
