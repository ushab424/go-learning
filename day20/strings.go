package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var greet string = "Привет"
	fmt.Println(len(greet))
	fmt.Println(utf8.RuneCountInString(greet))
	bytes := []byte(greet)
	fmt.Println(bytes)
	for _, value := range greet {
		fmt.Printf("%c;%d ", value, value)
	}
}
