package main

import "fmt"

func main() {
	const secret int = 7
	var a int
	guessed := false
	for i := 1; i <= 3; i++ {
		fmt.Print("Введите число которое загадала программа: ")
		fmt.Scan(&a)
		if a == secret {
			guessed = true
			fmt.Printf("Поздравляем! Вы угадали с %d раза!", i)
			break
		} else if a > secret {
			fmt.Println("Меньше")
		} else if a < secret {
			fmt.Println("Больше")
		}
	}
	if !guessed {
		fmt.Println("Попытки закончились!")
	}
}
