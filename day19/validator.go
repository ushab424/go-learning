package main

import "fmt"

type Validator interface {
	Validate(input string) bool
}

type MinLenght struct {
	Min int
}
type MaxLenght struct {
	Max int
}

func (mi *MinLenght) Validate(input string) bool {
	if len(input) >= mi.Min {
		return true
	} else {
		return false
	}
}
func (ma *MaxLenght) Validate(input string) bool {
	if len(input) <= ma.Max {
		return true
	} else {
		return false
	}
}

func checkAll(validators []Validator, input string) bool {
	for _, val := range validators {
		if !val.Validate(input) {
			return false
		}
	}
	return true
}

func main() {
	value := []Validator{&MaxLenght{10}, &MinLenght{3}}
	fmt.Println(checkAll(value, "hi"))
	fmt.Println(checkAll(value, "hello"))
	fmt.Println(checkAll(value, "verylongstring"))
}
