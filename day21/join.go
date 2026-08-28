package main

import (
	"strings"
)

func BuildSentence(words []string) string {
	words[0] = strings.ToUpper(string([]rune(words[0])[:1])) + string([]rune(words[0])[1:])
	return strings.Join(words, " ") + "."
}
