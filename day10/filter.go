package main

import "fmt"

func main() {
	var expensive []int
	prices := []int{150, 720, 300, 1500, 490, 800, 50, 999}
	for _, m := range prices {
		if m > 500 {
			expensive = append(expensive, m)
		}
	}
	fmt.Println(expensive)
	fmt.Println(len(expensive))
}
