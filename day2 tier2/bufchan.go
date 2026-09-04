package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	ch <- "a"
	ch <- "b"
	ch <- "c"
	for range 3 {
		fmt.Println(<-ch)
	}
}
