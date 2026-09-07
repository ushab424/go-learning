package main

import "fmt"

func main() {
	chen := make(chan int, 3)
	chen <- 1
	chen <- 2
	chen <- 3
	fmt.Println(<-chen)
	fmt.Println(<-chen)
	fmt.Println(<-chen)
}
