package main

import (
	"errors"
	"fmt"
)

func ReadConfig(path string) error {
	if path == "" {
		err := errors.New("empty file")
		return fmt.Errorf("ReadConfig: %w", err)
	} else {
		return nil
	}
}

func main() {
	conf := ReadConfig("")
	if conf != nil {
		fmt.Println(conf)
	} else {
		fmt.Println("cool")
	}

	conf2 := ReadConfig("ServerList")
	if conf2 != nil {
		fmt.Println(conf2)
	} else {
		fmt.Println("cool")
	}
}
