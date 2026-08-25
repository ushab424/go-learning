package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct{}

type Cat struct{}

type Cow struct{}

func (d *Dog) Sound() string {
	return "Woof"
}

func (c *Cat) Sound() string {
	return "Meow"
}

func (c *Cow) Sound() string {
	return "Moo"
}

func Choir(Animals []Animal) {
	for _, voice := range Animals {
		fmt.Println(voice.Sound())
	}
}

func main() {
	animals := []Animal{}
	animals = append(animals, &Dog{}, &Cat{}, &Cow{}, &Dog{}, &Cow{})
	Choir(animals)
}
