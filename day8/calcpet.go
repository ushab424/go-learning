package main

import "fmt"

func main() {
	var num1, num2, otvet int
	var operatsiya string
	story := [5]int{}
	for i := 1; i <= 5; i++ {
		fmt.Println("введите пример:")
		fmt.Scan(&num1, &operatsiya, &num2)
		switch operatsiya {
		case "+":
			slozh := plus(num1, num2)
			story[i-1] = slozh
			fmt.Println(slozh)
		case "-":
			vichetanie := vich(num1, num2)
			story[i-1] = vichetanie
			fmt.Println(vichetanie)
		case "*":
			umnozhenie := umnozh(num1, num2)
			story[i-1] = umnozhenie
			fmt.Println(umnozhenie)
		case "/":
			delenie := dele(num1, num2)
			story[i-1] = delenie
			fmt.Println(delenie)
		default:
			fmt.Println("Неизвестная операция!")
		}
		fmt.Println("еще?(1-да, 0-нет)")
		fmt.Scan(&otvet)
		if otvet != 1 {
			break
		}
	}
	fmt.Println(story)
}
func plus(num1, num2 int) int {
	itog := num1 + num2
	return itog
}
func vich(num1, num2 int) int {
	itog := num1 - num2
	return itog
}
func umnozh(num1, num2 int) int {
	itog := num1 * num2
	return itog
}
func dele(num1, num2 int) int {
	var itog int
	if num2 == 0 {
		itog = 0
	} else {
		itog = num1 / num2
	}
	return itog
}
