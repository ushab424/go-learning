package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I like Java, Java is cool."
	fmt.Println(strings.ReplaceAll(str, "Java", "Go"))
}
