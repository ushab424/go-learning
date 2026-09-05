package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func gr1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop")
			return
		default:
			fmt.Println("work")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go gr1(ctx, &wg)
	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
}
