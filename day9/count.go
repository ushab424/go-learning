package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Go is good. Go is great. Go is growing"
	str1 := " hello "
	fmt.Println(strings.Count(str, "Go"))
	fmt.Println(strings.TrimSpace(str1))
}
