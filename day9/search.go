package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Go is a great programming language."
	fmt.Println(strings.Contains(str, "great"))
	fmt.Println(strings.Index(str, "programming"))
}
