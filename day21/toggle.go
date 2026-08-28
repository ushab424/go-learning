package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Toggle(s string) string {
	letters := []rune(s)
	var result strings.Builder
	for _, val := range letters {
		if unicode.IsUpper(val) {
			result.WriteRune(unicode.ToLower(val))
		} else {
			result.WriteRune(unicode.ToUpper(val))
		}
	}
	return result.String()
}
func main() {
	fmt.Println(Toggle("Hello"))
}
