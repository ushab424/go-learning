package main

import "fmt"

func main() {
	var num [5]int = [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(num); i++ {
		fmt.Println(num[i])
	}
}
