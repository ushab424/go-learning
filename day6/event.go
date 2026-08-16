package main

import "fmt"

func main() {
	num := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(num); i++ {
		if num[i]%2 != 0 {
			continue
		}
		fmt.Println(num[i])
	}
}
