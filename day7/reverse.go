package main

import "fmt"

func main() {
	mas := [5]int{1, 2, 3, 4, 5}
	arr := reverse(mas)
	fmt.Println(arr)
}
func reverse(mas [5]int) (arr [5]int) {
	for i := 0; i < len(mas); i++ {
		arr[i] = mas[len(mas)-1-i]
	}
	return arr
}
