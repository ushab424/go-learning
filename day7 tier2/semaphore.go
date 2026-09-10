package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	tickets chan struct{}
}

func NewSemaphore(ticketNumber int) Semaphore {
	return Semaphore{
		tickets: make(chan struct{}, ticketNumber),
	}
}

func (s *Semaphore) Acquire() {
	s.tickets <- struct{}{}
}
func (s *Semaphore) Release() {
	<-s.tickets
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)

	semaphore := NewSemaphore(5)
	for i := 0; i < 6; i++ {
		semaphore.Acquire()
		go func() {
			defer func() {
				wg.Done()
				semaphore.Release()
			}()

			fmt.Println("working...")
			time.Sleep(2 * time.Second)
			fmt.Println("exiting...")
		}()
	}

	wg.Wait()
}
