package main

import (
	"errors"
	"fmt"
)

func LoadConfig(path, format string) (string, error) {
	if path == "" {
		err := errors.New("Пустой путь")
		return "", err
	}
	if format != "json" && format != "yaml" {
		err := errors.New("неизвестный формат")
		return "", err
	}
	return "loaded: " + path, nil
}
func main() {
	conf1, err := LoadConfig("", "json")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(conf1)
	}
	conf2, err2 := LoadConfig("d/c/p/h", "fijg")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(conf2)
	}
	conf3, err3 := LoadConfig("d/c/p/h", "json")
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(conf3)
	}
}
