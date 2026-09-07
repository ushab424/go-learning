package main

import "fmt"

func one(cha1 chan int) {
	for i := 1; i < 4; i++ {
		cha1 <- i
	}
	close(cha1)
}

func ten(cha2 chan int) {
	cha2 <- 10
	cha2 <- 20
	cha2 <- 30
	close(cha2)
}

func main() {
	cha1 := make(chan int)
	cha2 := make(chan int)
	go one(cha1)
	go ten(cha2)
	for val := range cha1 {
		fmt.Println(val)
	}
	for val := range cha2 {
		fmt.Println(val)
	}
}
