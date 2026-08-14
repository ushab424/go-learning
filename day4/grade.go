package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите ваш балл:")
	fmt.Scan(&a)
	if a >= 90 {
		fmt.Println("Отлично!")
	} else if a >= 70 && a <= 89 {
		fmt.Println("Хорошо")
	} else if a >= 50 && a <= 69 {
		fmt.Println("Удовлетворительно")
	} else {
		fmt.Println("Неудовлетворительно")
	}
}
