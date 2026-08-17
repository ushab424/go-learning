package main

import "fmt"

func main() {
	var b int
	fmt.Println("введите число:")
	fmt.Scan(&b)
	chet := rast(b)
	fmt.Println(chet)
}
func rast(b int) (z bool) {
	if b%2 == 0 {
		z = true
	}
	return z
}
