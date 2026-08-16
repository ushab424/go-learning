package main

import "fmt"

func main() {
	num := [5]int{3, 7, 2, 8, 5}
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	fmt.Println(sum)
}
