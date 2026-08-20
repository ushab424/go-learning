package main

import "fmt"

func main() {
	prices := []int{100, 250, 50, 300, 75}
	for i, dollars := range prices {
		fmt.Println(i, dollars)
	}
}
