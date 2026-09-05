package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func gr2(ctx context.Context, wg *sync.WaitGroup) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go gr2(ctx, &wg)
	wg.Wait()
}
