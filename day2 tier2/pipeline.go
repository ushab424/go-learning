package main

import "fmt"

func generate(n int, out chan int) {
	for i := 1; i <= n; i++ {
		out <- i
	}
	close(out)
}

func square(in chan int, out chan int) {
	for val := range in {
		out <- val * val
	}
	close(out)
}

func main() {
	in := make(chan int)
	out := make(chan int)
	go generate(5, in)
	go square(in, out)
	for val := range out {
		fmt.Println(val)
	}
}
