package main

import (
	"fmt"
)

func Reverse(s string) string {
	word := []rune(s)
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		// два указателя идут навстречу друг к другу
		word[i], word[j] = word[j], word[i]
	}
	return string(word)
}

func Max(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("Error: nums is empty")
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max, nil
}

func Contains(slice []string, target string) bool {
	for _, str := range slice {
		if str == target {
			return true
		}
	}
	return false
}
