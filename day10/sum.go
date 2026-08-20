package main

import "fmt"

func main() {
	var a int
	num := []int{10, 20, 30, 40, 50}
	for _, v := range num {
		a += v
	}
	fmt.Println(a)
}
