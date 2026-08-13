package main

import "fmt"

func main() {
	const rate float64 = 92.5
	var a float64
	fmt.Println("Введите сумму в рублях:")
	fmt.Scan(&a)
	fmt.Printf("%.2f руб = %.2f $\n", a, a/rate)
}
