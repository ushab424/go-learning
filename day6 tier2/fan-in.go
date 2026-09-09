package main

import (
	"fmt"
	"sync"
)

func fanIn(channels ...<-chan string) <-chan string {
	// принимаем входное множество string каналов
	wg := sync.WaitGroup{}
	// вейтгруппа для каждого канала из множества
	wg.Add(len(channels))
	out := make(chan string)
	// создаем результирующий канал
	for _, channel := range channels {
		go func() {
			// фоновая горутина для каждого канала из множества
			defer wg.Done()
			for value := range channel {
				// вычитывает данные из каждого канала
				out <- value
				// записывает данные в результирующий канал
			}
		}()
		// запускаем фоновую джобу которая асинхронно процессит
	}

	go func() {
		// ждем пока допишут горутины (фоново) в отдельной (фоновой) горутине
		wg.Wait()
		close(out)
	}()

	return out
	// сразу возвращаем результирующий канал и в фоне начинаем процессить
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	// создаем множество каналов
	go func() {
		// фоново закрываем каналы в конце
		defer func() {
			close(ch1)
			close(ch2)
			close(ch3)
		}()

		for i := 0; i < 100; i += 3 {
			// пишем в каналы
			ch1 <- "robot"
			ch2 <- "tobot"
			ch3 <- "auto"
		}
	}()

	for val := range fanIn(ch1, ch2, ch3) {
		// рейнж идет до тех пор пока канал не закроется
		fmt.Println(val)
	}
}
