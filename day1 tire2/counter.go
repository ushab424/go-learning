package main

import (
	"fmt"
	"sync"
)

func addCounter(counter *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		*counter++
		mu.Unlock()
	}
}

func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go addCounter(&counter, &wg, &mu)
	}
	wg.Wait()
	fmt.Println(counter)
}
