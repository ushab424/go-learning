package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Введите два числа:")
	fmt.Scan(&a, &b)
	fmt.Printf("%.d + %.d = %.d\n", a, b, a+b) //в кавычках шаблон вывода, дальше переменные
	fmt.Printf("%.d - %.d = %.d\n", a, b, a-b)
	fmt.Printf("%.d * %.d = %.d\n", a, b, a*b)
	fmt.Printf("%.d / %.d = %d\n", a, b, a/b)
	fmt.Printf("%.d %% %.d = %.d\n", a, b, a%b)
}
