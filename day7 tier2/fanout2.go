package main

import (
	"fmt"
	"sync"
)

func sqrt(id int, inputCH <-chan int, resCH chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range inputCH {
		resCH <- val * val
	}
}

func main() {
	inputCH := make(chan int)
	resCH := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(4)
	for i := 0; i < 4; i++ {
		go sqrt(i, inputCH, resCH, &wg)
	}
	go func() {
		for i := 1; i <= 20; i++ {
			inputCH <- i
		}
		close(inputCH)
	}()

	go func() {
		wg.Wait()
		close(resCH)
	}()

	for val := range resCH {
		fmt.Println(val)
	}
}
