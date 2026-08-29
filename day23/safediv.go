package main

import (
	"fmt"
)

func SafeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	result = a / b
	return result, nil
}
func main() {
	result, err := SafeDivide(10, 2)
	fmt.Println(result, err)
	result2, err2 := SafeDivide(10, 0)
	fmt.Println(result2, err2)
}
