package main

import (
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func (v *ValidationError) Error() string {
	return v.Field + ": " + v.Message
}

func ValidateName(name string) error {
	if name == "" {
		return &ValidationError{Field: "name", Message: "cant be empty"}
	}
	return nil
}

func main() {
	err := ValidateName("")
	if err != nil {
		fmt.Println(err)
	}
	err = ValidateName("Ivan")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("good")
	}
}
