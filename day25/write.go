package main

import (
	"fmt"
	"os"
)

func main() {
	greetings, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greetings, "ok")
	}
	defer greetings.Close()
	greetings.WriteString("Привет, Go!")
}
