package main

import (
	"strings"
)

func CountWords(s string) int {
	// comeback quantity words
	return len(strings.Fields(s))
}

func Filter(nums []int, fn func(int) bool) []int {
	// returns a new slice containing only the elements for which fn returned true.
	result := []int{}
	for i := 0; i < len(nums); i++ {
		if fn(nums[i]) {
			result = append(result, nums[i])
		}
	}
	return result
}

func Merge(a, b map[string]int) map[string]int {
	result := make(map[string]int)
	for k, v := range a {
		result[k] = v
	}
	for k, v := range b {
		result[k] = v
	}
	return result
}
