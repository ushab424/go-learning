package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("work1")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("worker 1 done")
}
func Worker2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("work2")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("worker 2 done")
}
func Worker3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("work3")
	time.Sleep(700 * time.Millisecond)
	fmt.Println("worker 3 done")
}
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go Worker1(&wg)
	go Worker2(&wg)
	go Worker3(&wg)
	wg.Wait()
	fmt.Println("Exit.")
}
