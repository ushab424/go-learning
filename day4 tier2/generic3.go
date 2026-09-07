package main

import "fmt"

type Pair[T any] struct {
	First  T
	Second T
}

func (p Pair[T]) Print() {
	fmt.Println(p.First, p.Second)
}

func main() {
	p1 := Pair[int]{1, 2}
	p1.Print() // 1 2

	p2 := Pair[string]{"hello", "world"}
	p2.Print() // hello world

}
