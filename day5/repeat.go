package main

import (
	"fmt"
	"strings"
)

func Repeat(s string, n int) string {
	var result strings.Builder
	for i := 0; i <= n; i++ {
		if i < n {
			result.WriteString(" ")
		}
		result.WriteString(s)
	}
	return result.String()
}

func main() {
	fmt.Println(Repeat("Hi", 5))
}
