package main

import "fmt"

func main() {
	var a float64
	const b float64 = 0.10
	const c float64 = 0.15
	fmt.Println("Сумма покупки:")
	fmt.Scan(&a)
	if a > 10000 {
		fmt.Printf("%.2f\n", a-(a*c))
	} else if a > 5000 {
		fmt.Printf("%.2f\n", a-(a*b))
	} else {
		fmt.Printf("%.2f\n", a)
	}
}
