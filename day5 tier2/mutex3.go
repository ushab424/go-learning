package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	data map[string]int
	sync.RWMutex
}

func (s *SafeMap) Read(key string) int {
	s.RLock()
	result := s.data[key]
	s.RUnlock()
	return result
}
func (s *SafeMap) Write(key string, val int) {
	s.Lock()
	s.data[key] = val
	s.Unlock()
}

func main() {
	data1 := &SafeMap{data: make(map[string]int)}
	wg := sync.WaitGroup{}
	wg.Add(20)
	for range 10 {
		go func() {
			defer wg.Done()
			data1.Write("Steal", 10)
		}()
	}
	for range 10 {
		go func() {
			defer wg.Done()
			data1.Read("Steal")
		}()
	}
	wg.Wait()
	fmt.Println(data1.Read("Steal"))
}
