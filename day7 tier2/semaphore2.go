package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	sem := make(chan struct{}, 3)
	// буфер канал, это и есть семафор. struct{} занимает 0 байт
	for i := 0; i < 20; i++ {
		wg.Add(1)
		// говорим что будет еще одна горутина
		sem <- struct{}{}
		/*
			занимаем слот в семафоре. Если 3 слота уже заняты, main блокируется здесь
			и ждет пока какая то горутина освободит слот
			поэтому одновременно работают 3
		*/
		go func(id int) {
			// запускаем горутину и передаем i как параметр что бы не было data race
			defer wg.Done()
			// уменьшить счетчик WaitGroup
			defer func() { <-sem }()
			// при завершении горутины читаем из семафораб освобождая слот для следующей
			fmt.Println("start", id)
			time.Sleep(1 * time.Second)
			fmt.Println("done", id)
		}(i)
		// передаем значение в горутину
	}

	wg.Wait()
}
