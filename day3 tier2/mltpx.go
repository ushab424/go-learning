package main

import (
	"fmt"
	"time"
)

func grfirst(ch1 chan string) {
	for range 3 {
		ch1 <- "gr1 work"
		time.Sleep(100 * time.Millisecond)
	}
}
func grsecond(ch2 chan string) {
	for range 3 {
		ch2 <- "gr2 work"
		time.Sleep(200 * time.Millisecond)
	}
}
func gr3(ch3 chan string) {
	for range 3 {
		ch3 <- "gr3 work"
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	go grfirst(ch1)
	go grsecond(ch2)
	go gr3(ch3)
	for range 9 {
		select {
		case val := <-ch1:
			fmt.Println(val)
		case val := <-ch2:
			fmt.Println(val)
		case val := <-ch3:
			fmt.Println(val)
		}
	}
}
