package main

import "fmt"

func noEscape() int {
	x := 42
	return x
}
func escapes() *int {
	x := 42
	return &x
}

func main() {
	x1 := noEscape()
	x2 := escapes()
	fmt.Println(x1)
	fmt.Println(*x2)
}
