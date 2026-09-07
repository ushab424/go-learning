package main

import "fmt"

func max[T int | float64](a, b T) T {
	var result T
	if a > b {
		result = a
	} else {
		result = b
	}
	return result
}

func main() {
	result := max(5, 7.15)
	fmt.Println(result)
}
