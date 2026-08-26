package main

import (
	"fmt"
	"sort"
)

type Sorter interface {
	Sort(nums []int) []int
}

type Asc struct{}
type Desc struct{}

func (a *Asc) Sort(nums []int) []int {
	sort.Ints(nums)
	return nums
}

func (d *Desc) Sort(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return nums
}

func PrintSorted(s Sorter, nums []int) {
	for _, val := range s.Sort(nums) {
		fmt.Println(val)
	}
}

func main() {
	numbers := []int{5, 2, 8, 1, 9}
	PrintSorted(&Asc{}, numbers)
	PrintSorted(&Desc{}, numbers)
}
