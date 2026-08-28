package main

import (
	"fmt"
	"strconv"
)

func SumFromStrings(a, b string) string {
	first, _ := strconv.Atoi(a)
	last, _ := strconv.Atoi(b)
	result := strconv.Itoa(first + last)
	return result
}

func main() {
	fmt.Println(SumFromStrings("5", "3"))
}
