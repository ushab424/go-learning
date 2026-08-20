package main

import "fmt"

func main() {
	colors := []string{"red", "green", "blue", "yellow", "pink"}
	colors = append(colors[:2], colors[3:]...)
	fmt.Println(colors)
}
