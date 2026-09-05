package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	select {
	case str := <-ch:
		fmt.Println(str)
	default:
		fmt.Println("channel empty")
	}
	go func() {
		ch <- "hello world"
	}()
	time.Sleep(100 * time.Millisecond)
	select {
	case str := <-ch:
		fmt.Println(str)
	default:
		fmt.Println("channel empty")
	}
}
