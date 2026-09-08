package main

import (
	"fmt"
	"sync"
)

type BancAccount struct {
	Balance int
	sync.Mutex
}

func (b *BancAccount) Deposit(amount int) {
	b.Mutex.Lock()
	b.Balance += amount
	b.Mutex.Unlock()
}
func (b *BancAccount) Withdraw(amount int) {
	b.Mutex.Lock()
	b.Balance -= amount
	b.Mutex.Unlock()
}

func main() {
	acc := &BancAccount{}
	wg := sync.WaitGroup{}
	wg.Add(100)
	for range 50 {
		go func() {
			defer wg.Done()
			acc.Deposit(100)
		}()
	}
	for range 50 {
		go func() {
			defer wg.Done()
			acc.Withdraw(100)
		}()
	}
	wg.Wait()
	fmt.Println(acc.Balance)
}
