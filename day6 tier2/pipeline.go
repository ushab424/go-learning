package main

import (
	"fmt"
)

func generate[T any](values ...T) <-chan T {
	outputCh := make(chan T)

	go func() {
		defer close(outputCh)
		for _, value := range values {
			outputCh <- value
		}
	}()

	return outputCh
}

func process[T any](inputCh <-chan T, action func(T) T) <-chan T {
	outputCh := make(chan T)

	go func() {
		defer close(outputCh)
		for value := range inputCh {
			outputCh <- action(value)
		}
	}()

	return outputCh
}

func main() {
	values := []int{1, 2, 3, 4, 5}
	mul := func(value int) int {
		return value * value
	}
	for value := range process(generate(values...), mul) {
		fmt.Println(value)
	}
}
