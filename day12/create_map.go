package main

import "fmt"

func main() {
	ages := map[string]int{
		"alice":   25,
		"bob":     30,
		"charlie": 22,
	}
	fmt.Println(ages)
	fmt.Println(ages["bob"])
}
