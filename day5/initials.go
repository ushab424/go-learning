package main

import (
	"fmt"
	"strings"
)

func GetInitials(name string) string {
	words := strings.Split(name, " ")
	var result strings.Builder
	for _, word := range words {
		runes := []rune(word)
		result.WriteRune(runes[0])
		result.WriteString(".")
	}
	return result.String()
}

func main() {
	fmt.Println(GetInitials("Иван Петров"))
	fmt.Println(GetInitials("Go Programming Language"))
}
