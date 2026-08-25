package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Widht, Height float64
}

func (c *Circle) Area() float64 {
	return 3.14 * (c.Radius * 2)
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Widht
}

func (c *Circle) Perimeter() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Height * r.Widht)
}

func PrintInfo(s Shape) {
	fmt.Println(s.Area(), s.Perimeter())
}

func main() {
	circle := Circle{Radius: 5.60}
	rectangle := Rectangle{Widht: 5, Height: 5.5}
	PrintInfo(&circle)
	PrintInfo(&rectangle)
}
