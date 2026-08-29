package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("не найден")

func FindUser(name string) error {
	if name == "" {
		return fmt.Errorf("FindUser: %w", ErrNotFound)
	} else {
		return nil
	}
}

func main() {
	err := FindUser("")
	if errors.Is(err, ErrNotFound) {
		fmt.Println("User not Found")
	}
}
