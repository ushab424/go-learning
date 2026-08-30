package main

import (
	"fmt"
	"strings"
)

func Censor(text string, banned []string) string {
	for _, word := range banned {
		text = strings.Replace(text, word, "**", -1)
	}
	return text
}

func main() {
	fmt.Println(Censor("Go is awesome and go is fun", []string{"awesome", "fun"}))
}
