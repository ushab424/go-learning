package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	nums := []int{10, -3, 7, -5, 12}
	// создаем слайс чисел
	ctx := context.Background()
	// корневой контекст
	g, qctx := errgroup.WithContext(ctx)
	// создаем errgroup с контекстом
	g.SetLimit(2)
	// ограничиваем одновременное количество выполняемых горутин до 2-х
	for _, num := range nums {
		// крутим слайс
		g.Go(func() error {
			// создаем горутину
			select {
			case <-ctx.Done():
				return qctx.Err()
			default:
				// если контекст отменен то выходим, в обратном случае идем дальше
			}
			if num < 0 {
				return fmt.Errorf("num %d is negative", num)
				// если число отрицательное, то возвращаем ошибку и выходим
			}
			fmt.Printf("%d: ok\n", num)
			return nil
			// если все ок, то печатаем и идем дальше
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
		// ждем горутины и в случае ошибки печатаем ее
	} else {
		fmt.Println("all done")
		// если ошибок нет, то печатаем что все ок
	}
}
