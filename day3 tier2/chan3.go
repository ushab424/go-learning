package main

import "fmt"

func ints(ch3 chan int) {
	for i := 1; i < 6; i++ {
		ch3 <- i
	}
	close(ch3)
}

func main() {
	ch3 := make(chan int)
	go ints(ch3)
	for val := range ch3 {
		fmt.Println(val)
	}
}
