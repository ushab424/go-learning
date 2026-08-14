package main

import "fmt"

func main() {
	var name string
	var age int
	var balance float64
	const sale float64 = 0.10
	fmt.Println("Введите ваше имя, возраст и баланс счета:")
	fmt.Scan(&name, &age, &balance)
	if age >= 18 {
		if balance > 1000 {
			fmt.Printf("Добро пожаловать, %s! Ваш статус: VIP. Баланс после скидки 10%%: %.2f\n", name, balance+(balance*sale))
		} else if balance >= 1 {
			fmt.Printf("Добро пожаловать, %s! Ваш статус: Обычный. Баланс: %.2f\n", name, balance)
		} else if balance <= 0 {
			fmt.Printf("Пополните счет!")
		}
	} else {
		fmt.Printf("Регистрация запрещена!")
	}
}
