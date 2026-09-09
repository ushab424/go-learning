package main

import "fmt"

func fib(n int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		if n <= 0 {
			output <- 0
		}
		a, b := 0, 1
		for i := 0; i < n; i++ {
			output <- a
			a, b = b, a+b
		}
	}()
	return output
}

func main() {
	ch := fib(5)
	for val := range ch {
		fmt.Println(val)
	}
}
