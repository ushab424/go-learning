package main

import (
	"errors"
	"fmt"
)

func ValidateProduct(name string, price float64, qty int) []error {
	err := []error{}
	if name == "" {
		err = append(err, errors.New("name empty"))
	}
	if price < 0 {
		err = append(err, errors.New("Price must be positive"))
	}
	if qty < 0 {
		err = append(err, errors.New("q-ty must be positive"))
	}
	return err
}
func main() {
	err := ValidateProduct("", -10, -3)
	for _, erro := range err {
		fmt.Println(erro)
	}
}
