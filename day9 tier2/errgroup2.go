package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	urls := []string{
		"https://api.good1.com",
		"https://api.good2.com",
		"https://api.bad.com",
		"https://api.good3.com",
		"https://api.good4.com",
	}
	ctx := context.Background()
	// корневой контекст
	g, qctx := errgroup.WithContext(ctx)
	// создаём группу и дочерний контекст. Когда любая горутина вернёт ошибку, qctx автоматически отменится.
	for _, url := range urls {
		// перебираем URL-ы, _ игнорирует индекс.
		g.Go(func() error {
			// запускаем горутины
			URL := url
			// копия для замыкания, чтобы каждая горутина видела свой URL.
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			// имитация сетевого запроса, случайная задержка 0-999ms.
			select {
			case <-qctx.Done():
				return qctx.Err()
			default:
				// если контекст уже отменён (другая горутина вернула ошибку), сразу выходим. default — если контекст жив, продолжаем.
			}
			if strings.Contains(URL, "bad") {
				return fmt.Errorf("url have bad word: %s", URL)
				// проверяем URL на "bad", возвращаем ошибку если нашли.
			}
			fmt.Println("ok:", URL)
			// если всё ок, печатаем URL.
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
		//  ждём все горутины, получаем первую ошибку.
	}

}
