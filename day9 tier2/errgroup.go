package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	// создаем корневой контекст
	g, _ := errgroup.WithContext(ctx)
	// создаём errgroup привязанную к ctx. g — это группа горутин.
	// Второй возвращаемый параметр — дочерний контекст который отменится при первой ошибке.
	for i := 1; i <= 5; i++ {
		n := i
		// переобьявляем во избежании data race
		g.Go(func() error {
			//  запускаем горутину через errgroup. Не нужно wg.Add и wg.Done — errgroup делает это сам. Функция должна вернуть error.
			if n%2 == 0 {
				return fmt.Errorf("num %d is even", n)
				// если чётное, возвращаем ошибку. errgroup запомнит первую ошибку.
			}
			fmt.Println(n, "ok")
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		// ждём завершения всех горутин. Wait() возвращает первую ошибку из всех горутин (или nil если ошибок не было).
		fmt.Println("Error:", err)
		//  печатаем ошибку.
	}
}
