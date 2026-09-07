package main

import "fmt"

func hello(ch1 chan string) {
	ch1 <- "hello"
}

func world(ch1 chan string) {
	ch1 <- "world"
}

func main() {
	ch1 := make(chan string)
	go hello(ch1)
	go world(ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
}
