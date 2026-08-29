package main

import (
	"errors"
	"fmt"
	"strconv"
)

func ParseID(input string) (int, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("ParseID: %w", err)
	}
	return id, nil
}

func CheckID(id int) error {
	if id <= 0 {
		err := errors.New("id must be positive")
		return fmt.Errorf("CheckID: %w", err)
	}
	return nil
}

func ProcessUser(input string) error {
	id, err := ParseID(input)
	if err != nil {
		return err
	}
	err = CheckID(id)
	if err != nil {
		return err
	}
	fmt.Println("User", id, "ok")
	return nil
}

func main() {
	us1 := ProcessUser("abc")
	if us1 != nil {
		fmt.Println(us1)
	}
	us2 := ProcessUser("-5")
	if us2 != nil {
		fmt.Println(us2)
	}
	us3 := ProcessUser("42")
	if us3 != nil {
		fmt.Println(us3)
	}
}
