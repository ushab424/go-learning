package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	a := Point{X: 1, Y: 2}
	b := Point{X: 1, Y: 2}
	c := Point{X: 3, Y: 4}
	fmt.Println(a == b, a == c)
}
