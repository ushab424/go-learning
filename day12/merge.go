package main

import "fmt"

func main() {
	m1 := map[string]int{
		"apple":  5,
		"banana": 3,
	}
	m2 := map[string]int{
		"banana": 7,
		"cherry": 2,
	}
	result := map[string]int{}

	for key, value := range m1 {
		result[key] += value
	}

	for key1, value1 := range m2 {
		result[key1] += value1
	}

	fmt.Println(result)
}
