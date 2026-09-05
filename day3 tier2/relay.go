package main

import (
	"fmt"
	"sync"
)

func fi(wg *sync.WaitGroup, f chan int, s chan int) {
	defer wg.Done()
	num := <-f
	s <- num * 2
}
func se(wg *sync.WaitGroup, s chan int, t chan int) {
	defer wg.Done()
	num := <-s
	t <- num + 10
}
func th(wg *sync.WaitGroup, t chan int) {
	defer wg.Done()
	fmt.Println(<-t)
}

func main() {
	f := make(chan int)
	s := make(chan int)
	t := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(3)
	go fi(&wg, f, s)
	go se(&wg, s, t)
	go th(&wg, t)
	f <- 5
	wg.Wait()
}
