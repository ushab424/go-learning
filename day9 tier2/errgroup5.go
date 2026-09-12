package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	files := []string{"a.txt", "b.txt", "error.txt", "c.txt", "d.txt", "error2.txt", "e.txt", "f.txt"}
	// есть слайс файлов на проверку
	ctx := context.Background()
	// создаем корневой контекст
	g, qctx := errgroup.WithContext(ctx)
	// создаем новую errgroup c контекстом отмены
	g.SetLimit(3)
	// ограничитель одновременных горутин
	for _, file := range files {
		// крутим слайс файлов
		g.Go(func() error {
			// создаем горутину
			time.Sleep(time.Millisecond * 300)
			// имитация загрузки
			select {
			case <-qctx.Done():
				return qctx.Err()
			default:
				// // если контекст отменен то выходим, в обратном случае идем дальше
			}
			if strings.Contains(file, "error") {
				return fmt.Errorf("file %s have our word", file)
				// // если имя файла содержит "error", то возвращаем ошибку и выходим
			}
			fmt.Println("file", file, "is ok!")
			return nil
			// если все ок, то печатаем и идем дальше
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
		// ждем горутины и в случае ошибки печатаем ее
	} else {
		fmt.Println("all files good!")
		// если ошибок нет, то печатаем что все ок
	}
}
