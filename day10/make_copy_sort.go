package main

import (
	"fmt"
	"sort"
)

func main() {
	prices := []int{50, 10, 40, 20, 30}
	pricessort := make([]int, 5)
	copy(pricessort, prices)
	sort.Ints(pricessort)
	fmt.Println(prices)
	fmt.Println(pricessort)
}
