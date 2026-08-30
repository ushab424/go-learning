package main

import (
	"fmt"
)

func Truncate(s string, max int) string {
	text := []rune(s)
	if len(text) > max {
		return string(text[:max]) + "..."
	}
	return s
}

func main() {
	fmt.Println(Truncate("Golang is very awesome language", 6))
}
