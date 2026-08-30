package main

import "fmt"

func FirstAndLast(s string) string {
	input := []rune(s)
	first := input[0]
	last := input[len(input)-1]
	return string(first) + "-" + string(last)
}

func main() {
	fmt.Println(FirstAndLast("Hello"))
}
