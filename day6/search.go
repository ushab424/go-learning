package main

import "fmt"

func main() {
	var a int
	num := [7]int{4, 8, 15, 16, 23, 42, 99}
	search := false
	fmt.Println("Введите число:")
	fmt.Scan(&a)
	for i := 0; i < len(num); i++ {
		if a == num[i] {
			fmt.Printf("Найдено на позиции %d", i)
			search = true
			break
		}
	}
	if !search {
		fmt.Println("Не обнаружено")
	}
}
