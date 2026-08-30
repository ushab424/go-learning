package main

import (
	"fmt"
	"strings"
)

func ReverseString(s string) string {
	reverse := []rune(s)
	var result strings.Builder
	for i := len(reverse) - 1; i >= 0; i-- {
		result.WriteRune(reverse[i])
	}
	return result.String()
}

func main() {
	fmt.Println(ReverseString("Привет"))
}
