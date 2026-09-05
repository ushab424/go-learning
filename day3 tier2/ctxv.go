package main

import (
	"context"
	"fmt"
	"sync"
)

func grr(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	user := ctx.Value("user")
	fmt.Println(user)
}

func main() {
	ctx := context.WithValue(context.Background(), "user", "admin")
	var wg sync.WaitGroup
	wg.Add(1)
	go grr(ctx, &wg)
	wg.Wait()
}
