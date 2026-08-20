package main

import "fmt"

func main() {
	codes := []int{200, 404, 200, 500, 200, 404, 500, 500, 200}
	result := make(map[int]int)
	for _, value := range codes {
		result[value]++
	}
	for code, count := range result {
		fmt.Println(code, count)
	}
}
