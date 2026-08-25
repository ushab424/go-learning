package main

import "fmt"

type Operation interface {
	Execute(a, b float64) float64
}

type Add struct{}

type Sub struct{}

type Mul struct{}

type Div struct{}

func (A *Add) Execute(a, b float64) float64 {
	return a + b
}

func (s *Sub) Execute(a, b float64) float64 {
	return a - b
}

func (m *Mul) Execute(a, b float64) float64 {
	return a * b
}
func (d *Div) Execute(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
func calc(op Operation, a, b float64) {
	fmt.Println(op.Execute(a, b))
}

func main() {
	num := Add{}
	num2 := Sub{}
	num3 := Mul{}
	num4 := Div{}
	calc(&num, 20, 4)
	calc(&num2, 20, 4)
	calc(&num3, 20, 4)
	calc(&num4, 20, 4)
}
