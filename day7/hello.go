package main

import "fmt"

func main() {
	var name string
	fmt.Println("Введите имя:")
	fmt.Scan(&name)
	hello(name)
}
func hello(name string) {
	fmt.Println("Привет,", name)
}
