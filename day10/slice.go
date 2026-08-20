package main

import "fmt"

func main() {
	id := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pack1 := id[:3]
	pack2 := id[7:]
	fmt.Println(pack1)
	fmt.Println(pack2)
}
