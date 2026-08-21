package main

import "fmt"

func main() {
	nums := []int{5, 3, 8, 1, 9, 2}
	var max int
	var max2 int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > max {
			max2 = max
			max = nums[i]
		} else if nums[i] > max2 {
			max2 = nums[i]
		}
	}
	fmt.Println(max2)
}
