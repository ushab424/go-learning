package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("worker %d: processing %d\n", id, job)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(3)
	go worker(1, jobs, results, &wg)
	go worker(2, jobs, results, &wg)
	go worker(3, jobs, results, &wg)
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()
	go func() {
		wg.Wait()
		close(results)
	}()

	for val := range results {
		fmt.Println(val)
	}
}
