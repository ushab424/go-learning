package main

import "fmt"

func Generate(n int, ch chan int) {
	for i := 1; i <= n; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go Generate(10, ch)
	for val := range ch {
		fmt.Println(val)
	}
}
