package main

import (
	"fmt"
	"strings"
)

type Formatter interface {
	Format(Data string) string
}

type Upper struct{}
type Wrapper struct {
	Prefix string
	Suffix string
}

func (u *Upper) Format(Data string) string {
	return strings.ToUpper(Data)
}
func (w *Wrapper) Format(Data string) string {
	return w.Prefix + Data + w.Suffix
}

func applyAll(formatters []Formatter, Data string) string {
	for _, f := range formatters {
		Data = f.Format(Data)
	}
	return Data
}

func main() {
	greetings := []Formatter{&Upper{}, &Wrapper{Prefix: "[", Suffix: "]"}}
	fmt.Println(applyAll(greetings, "hello"))
}
