package main

import (
	"fmt"
	"strconv"
)

func ConvertAll(inputs []string) ([]int, []error) {
	output := []int{}
	er := []error{}
	for _, word := range inputs {
		num, err := strconv.Atoi(word)
		if err != nil {
			er = append(er, err)
		} else {
			output = append(output, num)
		}
	}
	return output, er
}
func main() {
	conv, err := ConvertAll([]string{"10", "abc", "5", "xyz"})
	fmt.Println(conv, err)
}
