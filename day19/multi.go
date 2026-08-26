package main

import "fmt"

type Stringer interface {
	String() string
}

type Sizer interface {
	Size() int
}

type Message struct {
	Text string
}

func (m *Message) String() string {
	return m.Text
}

func (m *Message) Size() int {
	return len(m.Text)
}

func PrntStr(s Stringer) {
	fmt.Println(s.String())
}

func PrntSz(s Sizer) {
	fmt.Println(s.Size())
}

func main() {
	Messages := Message{"Hello"}
	PrntStr(&Messages)
	PrntSz(&Messages)
}
