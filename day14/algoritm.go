package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(nums); i++ {
		if nums[i] < nums[len(nums)-1] {
			nums[i] = nums[len(nums)-1]
		}
	}
	for i := 0; i >= nums[3]; i++ {
		if nums[i] > nums[3] {
			nums[i] = nums[3]
		}
	}
	for i := 3; i >= len(nums)-1; i++ {
		if nums[i] > nums[len(nums)-1] {
			nums[i] = nums[len(nums)-1]
		}
	}
	fmt.Println(nums)
}
