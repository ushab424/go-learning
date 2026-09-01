package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for i := 1; i <= 5; i++ {
		_, err := file.WriteString(fmt.Sprintf("Строка %d\n", i))
		if err != nil {
			fmt.Println("Ошибка записи:", err)
			return
		}
	}
	file2, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file2.Close()
	scanner := bufio.NewScanner(file2)
	num := 1
	for scanner.Scan() {
		fmt.Printf("%d: %s\n", num, scanner.Text())
		num++
	}
}
