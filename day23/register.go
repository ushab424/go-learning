package main

import (
	"errors"
	"fmt"
	"strings"
)

func Register(name, email, password string) error {
	if name == "" {
		return errors.New("name necessarily")
	}
	if strings.Contains(email, "@") == false {
		return errors.New("invalid email")
	}
	pass := []rune(password)
	if len(pass) < 6 {
		return errors.New("password too little")
	}
	return nil
}
func main() {
	err := Register("", "emsail@ojdig89", "20849578647867")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Registration succsesfull")
	}
	err1 := Register("Ivan", "emsailojdig89", "20849578647867")
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("Registration succsesfull")
	}
	err2 := Register("Ivan", "emsail@ojdig89", "2080")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Registration succsesfull")
	}
}
