package main

import (
	"fmt"
	"sync"
)

func SplitChannel[T any](inputCh <-chan T, n int) []<-chan T {
	outputChs := make([]chan T, n)
	for i := 0; i < n; i++ {
		outputChs[i] = make(chan T)
	}
	// create the resulting channel

	go func() {
		// create a background processing job
		idx := 0
		for value := range inputCh {
			outputChs[idx] <- value
			idx = (idx + 1) % n
		}

		for _, ch := range outputChs {
			close(ch)
		}
	}()

	resultChs := make([]<-chan T, n)
	// we have <- channel
	for i := 0; i < n; i++ {
		resultChs[i] = outputChs[i]
	}

	return resultChs
	// return the resilting channel
}

func main() {
	channel := make(chan int)
	// create channel

	go func() {
		// this function writing in channel
		defer close(channel)
		// after we close here
		for i := 0; i < 10; i++ {
			channel <- i
		}
	}()

	channels := SplitChannel(channel, 2)
	// SplitChannel divide the stream that comes out of this function by 2
	// SplitChannel return slice channels

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		// reading first channel value
		defer wg.Done()
		for value := range channels[0] {
			fmt.Println("ch1: ", value)
		}
	}()

	go func() {
		// reading second channel value
		defer wg.Done()
		for value := range channels[1] {
			fmt.Println("ch2: ", value)
		}
	}()

	wg.Wait()

	// we have 1 channel => we have 2 channel
}
