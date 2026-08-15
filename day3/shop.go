package main

import "fmt"

func main() {
	const price float64 = 1500
	const sale float64 = 0.10
	var a float64
	fmt.Println("Введите количество товара:")
	fmt.Scan(&a)
	fmt.Printf("Количество: %.2f\n", a) // %.2f что бы деньги красивее выглядели
	fmt.Printf("Сумма без скидки: %.2f\n", a*price)
	fmt.Printf("Скидка: %.2f\n", (a*price)*sale)
	fmt.Printf("Итого: %.2f\n", (a*price)-((a*price)*sale))
}
