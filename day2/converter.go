package main

import "fmt"

func main() {
	var c int
	fmt.Printf("введите температуру цельсия")
	fmt.Scan(&c)
	fmt.Println(c*9/5 + 32)
}
