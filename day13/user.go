package main

import "fmt"

type user struct {
	name, email string
	age         int
}

func main() {
	tom := user{"Tom", "tomasbag@gmail.com", 20}
	genry := user{"Genry", "genryhill@yandex.ru", 35}
	fmt.Println(tom, genry)
	fmt.Println(genry.email)
}
