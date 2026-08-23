package main

import "fmt"

func Normalize(scores *[]int) {
	var max int
	for i := 0; i <= len(*scores)-1; i++ {
		if (*scores)[i] > max {
			max = (*scores)[i]
		}
	}
	for i := 0; i <= len(*scores)-1; i++ {
		(*scores)[i] = (*scores)[i] * 100 / max
	}
}

func main() {
	scores := []int{45, 90, 30, 60}
	Normalize(&scores)
	fmt.Println(scores)
}
