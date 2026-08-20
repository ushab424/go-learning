package main

import "fmt"

func main() {
	cart := map[string]int{
		"phone":  50000,
		"laptop": 120000,
		"mouse":  1500,
	}
	delete(cart, "mouse")
	for key, value := range cart {
		fmt.Println(key, value)
	}
}
