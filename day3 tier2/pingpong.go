package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func first(ctx context.Context, ping chan string, pong chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-ping:
			fmt.Println("first:", msg)
			pong <- "pong"
		}
	}
}
func second(ctx context.Context, ping chan string, pong chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-pong:
			fmt.Println("second:", msg)
			ping <- "ping"
		}
	}
}

func main() {
	ping := make(chan string, 1)
	pong := make(chan string, 1)
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	wg.Add(2)
	go first(ctx, ping, pong, &wg)
	go second(ctx, ping, pong, &wg)
	ping <- "ping"
	wg.Wait()
	fmt.Println("game over")
}
