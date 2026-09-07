package main

import "fmt"

func contains[T comparable](slice []T, val T) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(contains([]int{1, 2, 3}, 2))
	fmt.Println(contains([]string{"a", "b"}, "c"))

}
