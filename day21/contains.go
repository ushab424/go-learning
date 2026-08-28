package main

import (
	"fmt"
	"strings"
)

func FindWord(text, word string) string {
	result := strings.Contains(text, word)
	if !result {
		return "Не найдено"
	} else {
		return "Найдено"
	}
}

func main() {
	fmt.Println(FindWord("go is awesome", "go"))
	fmt.Println(FindWord("go is awesome", "rust"))

}
