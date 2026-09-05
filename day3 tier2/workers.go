package main

import (
	"fmt"
	"sync"
)

func worker(id int, tasks chan int, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	num := <-tasks
	fmt.Println("worker", id, ":got", num)
	done <- true
}

func main() {
	wg := sync.WaitGroup{}
	tasks := make(chan int)
	done := make(chan bool)
	wg.Add(3)
	go worker(1, tasks, done, &wg)
	go worker(2, tasks, done, &wg)
	go worker(3, tasks, done, &wg)
	tasks <- 10
	tasks <- 20
	tasks <- 30
	for range 3 {
		fmt.Println(<-done)
	}
	wg.Wait()
}
