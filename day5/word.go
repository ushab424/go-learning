package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	input := strings.Split(s, " ")
	result := map[string]int{}
	for _, words := range input {
		result[words]++
	}
	return result
}

func main() {
	fmt.Println(WordCount("go is go and go is fun"))
}
