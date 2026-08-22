package main

import "fmt"

type Rectangle struct {
	Widht float64
	Hight float64
}

func (A Rectangle) Area() {
	fmt.Println(A.Widht * A.Hight) //Метод площади
}

func (P Rectangle) Perimeter() {
	fmt.Println((P.Hight + P.Widht) * 2) //Метод периметра
}

func main() {
	RectangleVariable := Rectangle{Widht: 20, Hight: 5}
	RectangleVariable.Area()
	RectangleVariable.Perimeter()
}
