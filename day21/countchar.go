package main

import (
	"fmt"
)

func Between(s string, start, end int) string {
	text := []rune(s)
	result := text[start:end]
	return string(result)
}

func main() {
	fmt.Println(Between("Golang cool", 7, 11))
}
