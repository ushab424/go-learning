package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	// создаем корневой контекст
	g, _ := errgroup.WithContext(ctx)
	// создаем errgroup без контекста отмены
	g.SetLimit(3)
	// ограничиваем одновременное выполнение горутин
	for i := 1; i < 11; i++ {
		// создаем 10 горутин
		num := i
		// переобьявляем переменную
		g.Go(func() error {
			// создаем горутинку
			fmt.Println("start", num)
			// печатаем о начале работы горутины
			time.Sleep(time.Second * 1)
			// имитация работы
			fmt.Println("done:", num)
			// печатаем о завершении работы
			return nil
			// возвращаем нил ошибку
		})
	}
	g.Wait()
	// ждем окончания работы всех горутин
}
