package main

import (
	"fmt"
	"sync"
	"time"
)

func printA(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println("A:", i)
		time.Sleep(100 * time.Millisecond)
	}
}
func printB(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println("B:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printA(&wg)
	go printB(&wg)
	wg.Wait()
}
