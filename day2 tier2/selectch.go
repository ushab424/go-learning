package main

import (
	"fmt"
	"time"
)

func gr1(ch1 chan string) {
	time.Sleep(200 * time.Millisecond)
	ch1 <- "from gr1"
}

func gr2(ch2 chan string) {
	time.Sleep(500 * time.Millisecond)
	ch2 <- "from gr2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go gr1(ch1)
	go gr2(ch2)
	for i := 0; i < 2; i++ {
		select {
		case str := <-ch1:
			fmt.Println(str)
		case str := <-ch2:
			fmt.Println(str)
		}
	}
}
