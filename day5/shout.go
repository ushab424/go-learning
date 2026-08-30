package main

import (
	"fmt"
	"strings"
)

func Shout(s string) string {
	s = strings.ToUpper(s)
	return s + "!!!"
}

func main() {
	fmt.Println(Shout("hello"))
}
