package main

import (
	"fmt"
	"strings"
)

func HideNumber(phone string) string {
	phonenum := []rune(phone)
	lenmums := len(phonenum) - 4
	return strings.Repeat("*", lenmums) + string(phonenum[lenmums:])
}
func main() {
	fmt.Println(HideNumber("08938947846889"))
}
