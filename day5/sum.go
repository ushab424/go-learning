package main

import "fmt"

func main() {
	var a int
	sum := 0
	fmt.Println("Введите число")
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		sum += i
	}
	fmt.Println(sum)
}
