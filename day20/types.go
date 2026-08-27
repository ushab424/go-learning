package main

import "fmt"

func main() {
	var a int8 = 127
	var b uint8 = 255
	var c float32 = 3.14
	var d float64 = 3.141592653589793
	var e rune = 'Я'
	var f byte = 'A'
	fmt.Printf("%T %d\n", a, a)
	fmt.Printf("%T %d\n", b, b)
	fmt.Printf("%T %f\n", c, c)
	fmt.Printf("%T %f\n", d, d)
	fmt.Printf("%T %c %d\n", e, e, e)
	fmt.Printf("%T %c %d\n", f, f, f)
	a++
	fmt.Printf("a после a++: %d\n", a)
}
