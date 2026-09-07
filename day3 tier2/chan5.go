package main

import "fmt"

func num(channel chan int) {
	for i := 1; i < 6; i++ {
		channel <- i
	}
	close(channel)
}

func main() {
	channel := make(chan int)
	go num(channel)
	for {
		val, ok := <-channel
		if !ok {
			fmt.Println("exit")
			break
		}
		fmt.Println(val)
	}
}
