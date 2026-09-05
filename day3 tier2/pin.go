package main

import (
	"fmt"
	"sync"
)

func player1(wg *sync.WaitGroup, fir chan int, sec chan int) {
	defer wg.Done()
	for {
		num := <-fir
		if num >= 10 {
			fmt.Println("game over")
			close(sec)
			return
		}
		fmt.Println("player1:", num)
		sec <- (num + 1)
	}
}

func player2(wg *sync.WaitGroup, sec chan int, fir chan int) {
	defer wg.Done()
	for num := range sec {
		fmt.Println("player2:", num)
		fir <- num + 1
	}
}

func main() {
	fir := make(chan int)
	sec := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go player1(&wg, fir, sec)
	go player2(&wg, sec, fir)
	fir <- 0
	wg.Wait()
}
