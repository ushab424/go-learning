package main

import (
	"fmt"
	"strconv"
)

func Summary(name string, age int) string {
	agestr := strconv.Itoa(age)
	return "Возраст: " + agestr + " Имя: " + name
}

func main() {
	fmt.Println(Summary("Ivan", 25))
}
