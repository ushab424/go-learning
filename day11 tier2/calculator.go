package main

import (
	"fmt"
)

func Subtract(a, b int) int {
	return a - b
}
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("devide on 0")
	}
	return a / b, nil
}
