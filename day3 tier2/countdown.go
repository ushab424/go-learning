package main

import (
	"fmt"
	"time"
)

func timer(ch chan int) {
	for i := 5; i > 0; i-- {
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go timer(ch)
	for val := range ch {
		fmt.Println(val)
	}
	fmt.Println("GO!")
}
