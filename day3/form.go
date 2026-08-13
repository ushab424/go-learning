package main

import "fmt"

func main() {
	var name string
	var age int
	var salary float64
	const tax float64 = 0.13
	fmt.Println("Введите ваше имя, возраст и зарплату:")
	fmt.Scan(&name, &age, &salary)
	fmt.Printf("Имя: %s\n", name)
	fmt.Printf("Возраст: %d\n", age)
	fmt.Printf("Зарплата: %.2f\n", salary)
	fmt.Printf("Налог (13%%): %.2f\n", salary*tax)
	fmt.Printf("На руки: %.2f\n", salary-(salary*tax))
	fmt.Printf("Через 5 лет вам будет: %d\n", age+5)
}
