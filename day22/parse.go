package main

import (
	"errors"
	"fmt"
	"strconv"
)

func ParseAge(input string) (int, error) {
	nums, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	if nums < 0 {
		return 0, errors.New("Age is too little")
	}
	if nums > 150 {
		return 0, errors.New("Age is too much")
	}
	return nums, nil
}
func main() {
	output1, err := ParseAge("25")
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(output1)
	}
	output2, err := ParseAge("abc")
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(output2)
	}
	output3, err := ParseAge("-10")
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(output3)
	}
}
