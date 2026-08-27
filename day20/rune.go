package main

import (
	"fmt"
	"strings"
)

func main() {
	var greet string = "Golang - это круто"
	mgreet := []rune(greet)
	mgreet = mgreet[:6]
	result := string(mgreet[:6])
	fmt.Println(result)
	fmt.Println(strings.Replace(greet, "круто", "мощь", 1))
}
