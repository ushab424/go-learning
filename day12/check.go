package main

import "fmt"

func main() {
	users := map[string]bool{
		"admin": true,
		"guest": false,
	}
	value, ok := users["admin"]
	fmt.Println(value, ok)

	value2, ok2 := users["moderator"]
	fmt.Println(value2, ok2)
}
