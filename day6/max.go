package main

import "fmt"

func main() {
	sum := [5]int{11, 45, 3, 67, 24}
	max := sum[0]
	for i := 0; i < len(sum); i++ {
		if sum[i] > max {
			max = sum[i]
		}
	}
	fmt.Println(max)
}
