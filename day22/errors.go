package main

import (
	"errors"
	"fmt"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Деление на ноль")
	} else {
		return a / b, nil
	}
}

func main() {
	result, err := Divide(5, 3)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
	result1, err := Divide(5, 0)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result1)
	}

}
