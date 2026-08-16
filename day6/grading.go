package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите оценку:")
	fmt.Scan(&a)
	switch a {
	case 5:
		fmt.Println("Отлично!")
		fallthrough
	case 4:
		fmt.Println("Хорошо!")
		fallthrough
	case 3:
		fmt.Println("Зачет!")
	case 2, 1:
		fmt.Println("Незачет")
	}
}
