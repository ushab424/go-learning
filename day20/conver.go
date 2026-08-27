package main

import "fmt"

func main() {
	var x int = 10
	var y int = 3
	fmt.Println(x / y)
	fmt.Println(float64(x) / float64(y))
	var f float64 = 9.99
	fmt.Println(int(f))
	var big int64 = 300
	fmt.Println(int8(big))
}
