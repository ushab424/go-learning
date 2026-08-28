package main

import (
	"fmt"
	"strings"
)

func ParseTags(input string) []string {
	words := []string{}
	words = strings.Split(input, ",")
	for i, word := range words {
		words[i] = strings.TrimSpace(word)
	}
	return words
}

func main() {
	fmt.Println(ParseTags("go , rust , python "))
}
