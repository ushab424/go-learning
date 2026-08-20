package main

import "fmt"

func main() {
	var fruits []string
	fruits = append(fruits, "apple")
	fruits = append(fruits, "banana")
	fruits = append(fruits, "cherry")
	fmt.Println(fruits)
	fmt.Println(len(fruits))
}
