package main

import (
	"errors"
	"fmt"
)

func ValidateUser(name string, age int) []error {
	err := []error{}
	if name == "" {
		err = append(err, errors.New("Name empty"))
	}
	if age < 0 {
		err = append(err, errors.New("Age too little"))
	}
	if age > 150 {
		err = append(err, errors.New("Age too much"))
	}
	return err
}
func main() {
	errs := ValidateUser("Ivan", 25)
	for _, e := range errs {
		fmt.Println(e)
	}
}
