package main

import (
	"fmt"
	"os"
)

func main() {
	text, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(text))
}
