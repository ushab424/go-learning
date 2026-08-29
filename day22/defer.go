package main

import "fmt"

func main() {
	fmt.Println("начало")
	defer fmt.Println("отложено 1")
	defer fmt.Println("отложено 2")
	defer fmt.Println("отложено 3")
	fmt.Println("конец")
}
