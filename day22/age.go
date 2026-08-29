package main

import (
	"errors"
	"fmt"
)

func CheckAge(age int) (string, error) {
	if age < 0 {
		return "ERROR:", errors.New("Age uncorected")
	} else if age > 150 {
		return "", errors.New("Age uncorected")
	}
	return "Age corected", nil
}

func main() {
	output1, err := CheckAge(25)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(output1)
	}
	output2, err := CheckAge(-5)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(output2)
	}
	output3, err := CheckAge(200)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(output3)
	}
}
