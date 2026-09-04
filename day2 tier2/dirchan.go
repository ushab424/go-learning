package main

import "fmt"

func producer(ch chan<- int) {
	for i := 1; i < 6; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
