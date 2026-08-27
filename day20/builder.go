package main

import (
	"fmt"
	"strings"
)

func main() {
	before := []string{"Go", "is", "awesome"}
	var after strings.Builder
	for _, word := range before {
		after.WriteString(word)
		after.WriteString(" ")
	}
	fmt.Println(after.String())
}
