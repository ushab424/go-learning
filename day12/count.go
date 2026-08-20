package main

import "fmt"

func main() {
	words := []string{"go", "python", "go", "java", "go", "python"}
	result := make(map[string]int)
	for _, value := range words {
		result[value]++
	}
	for word, count := range result {
		fmt.Println(word, count)
	}
}
