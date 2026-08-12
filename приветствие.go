package main

import "fmt"

func main() {
	var name string
	var age int
	var a int = 10
	fmt.Print("Привет, как тебя зовут и сколько тебе лет?")
	fmt.Scan(&name, &age)
	fmt.Println("Привет,", name, "! Тебе", age, "лет, через 10 лет тебе будет", age+a)
}
