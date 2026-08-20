package main

import "fmt"

func main() {
	products := []string{"phone", "laptop", "phone", "tablet", "laptop", "phone", "mouse", "tablet"}
	result := make(map[string]int)
	for _, value := range products {
		result[value]++
	}
	for product, value := range result {
		if value > 1 {
			fmt.Println(product, value)
		}
	}
}
