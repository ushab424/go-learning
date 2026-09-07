package main

import "fmt"

func filter[T any](slice []T, fn func(T) bool) []T {
	result := []T{}
	for _, v := range slice {
		if fn(v) == true {
			result = append(result, v)
		}
	}
	return result
}
func main() {
	var1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(filter(var1, func(n int) bool {
		return n%2 == 0
	}))
	var2 := []string{"go", "hi", "rust", "ok"}
	fmt.Println(filter(var2, func(s string) bool {
		return len(s) > 2
	}))
}
