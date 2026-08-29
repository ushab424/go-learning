package main

import "fmt"

func safeDiv(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Поймали панику:", r)
		}
	}()
	fmt.Println(a / b)
}
func main() {
	safeDiv(10, 2)
	safeDiv(10, 0)
}
