package main

import (
	"fmt"
	"time"
)

func Producer1(ch chan string) {
	for i := 1; i < 4; i++ {
		ch <- fmt.Sprintf("p1: %d\n", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func Producer2(ch chan string) {
	for i := 1; i < 4; i++ {
		ch <- fmt.Sprintf("p2: %d\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	ch := make(chan string)
	go Producer1(ch)
	go Producer2(ch)
	for i := 0; i < 6; i++ {
		fmt.Println(<-ch)
	}
}
