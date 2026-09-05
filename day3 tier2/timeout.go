package main

import (
	"fmt"
	"time"
)

func gr1(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "done"
}

func main() {
	ch := make(chan string)
	go gr1(ch)
	select {
	case str := <-ch:
		fmt.Println(str)
	case <-time.After(5 * time.Second):
		fmt.Println("timeout")
	}
}
