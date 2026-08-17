package main

import "fmt"

func main() {
	numb := [5]int{3, 7, 2, 8, 5}
	sum := summary(numb)
	fmt.Println(sum)
}
func summary(numb [5]int) int {
	sum := 0
	for i := 0; i < len(numb); i++ {
		sum = numb[i] + sum
	}
	return sum
}
