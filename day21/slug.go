package main

import (
	"fmt"
	"strings"
)

func Slug(s string) string {
	result := strings.ToLower(s)
	result = strings.ReplaceAll(result, " ", "-")
	return result
}

func main() {
	fmt.Println(Slug("Go is awesome"))
}
