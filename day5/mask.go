package main

import (
	"fmt"
	"strings"
)

func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	runes := []rune(parts[0])
	first := string(runes[:2])
	return first + "****@" + parts[1]
}

func main() {
	fmt.Println(MaskEmail("ivan@gmail.com"))
	fmt.Println(MaskEmail("googleo@gmail.com"))
}
