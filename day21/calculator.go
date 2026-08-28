package main

import (
	"fmt"
	"strconv"
)

func CalcFromString(a, b, op string) string {
	first, _ := strconv.Atoi(a)
	last, _ := strconv.Atoi(b)
	var result int
	switch op {
	case "+":
		result = first + last
	case "-":
		result = first - last
	case "*":
		result = first * last
	case "/":
		if last == 0 {
			result = 0
		} else {
			result = first / last
		}
	}
	return strconv.Itoa(result)
}

func main() {
	fmt.Println(CalcFromString("5", "2", "+"))
}
