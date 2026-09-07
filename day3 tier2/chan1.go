package main

import (
	"fmt"
)

func sum(ch chan int) {
	ch <- 42
}

func main() {
	ch := make(chan int)
	go sum(ch)
	fmt.Println(<-ch)
}
