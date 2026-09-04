package main

import (
	"fmt"
)

func Sum(nums []int, ch chan int) {
	var result int
	for _, v := range nums {
		result += v
	}
	ch <- result
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	num1 := []int{}
	num2 := []int{}
	for i := 0; i < 5; i++ {
		num1 = append(num1, nums[i])
	}
	for i := 5; i < 10; i++ {
		num2 = append(num2, nums[i])
	}
	ch := make(chan int)
	go Sum(num1, ch)
	go Sum(num2, ch)
	result1 := <-ch
	result2 := <-ch
	fmt.Println(result1 + result2)
}
