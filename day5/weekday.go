package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите число от 1 до 7:")
	fmt.Scan(&a)
	if a >= 1 && a <= 7 {
		switch a {
		case 1:
			fmt.Println("Понедельник")
		case 2:
			fmt.Println("Вторник")
		case 3:
			fmt.Println("Среда")
		case 4:
			fmt.Println("Четверг")
		case 5:
			fmt.Println("Пятница")
		case 6:
			fmt.Println("Суббота")
		case 7:
			fmt.Println("Воскресенье")
		}
	} else {
		fmt.Println("Неизвестный день")
	}
}
