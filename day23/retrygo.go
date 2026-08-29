package main

import (
	"errors"
	"fmt"
)

func Retry(attemts int, fn func() error) error {
	var err error
	for i := 0; i < attemts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
	}
	return err
}

func main() {
	count := 0
	err := Retry(5, func() error {
		count++
		if count < 3 {
			fmt.Println("попытка", count, "— ошибка")
			return errors.New("не получилось")
		}
		fmt.Println("попытка", count, "— успех")
		return nil
	})
	if err != nil {
		fmt.Println("итог:", err)
	} else {
		fmt.Println("всё ок")
	}

}
