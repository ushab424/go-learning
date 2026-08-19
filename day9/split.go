package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "one:two:three:four:five"
	mas := strings.Split(str, ":")
	fmt.Println(mas)
	fmt.Println(mas[1])
}
