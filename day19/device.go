package main

import "fmt"

type Printer interface {
	Print() string
}
type Scanner interface {
	Scan() string
}
type Faxer interface {
	Fax() string
}

type AllInOne struct {
	Model string
}
type SimplePrinter struct {
	Model string
}

func (a *AllInOne) Print() string {
	return a.Model + " Printing"
}
func (a *AllInOne) Scan() string {
	return a.Model + " Scaning"
}
func (a *AllInOne) Fax() string {
	return a.Model + " Faxing"
}
func (s *SimplePrinter) Print() string {
	return s.Model + " Printing"
}

func Check(v interface{}) {
	if p, ok := v.(Printer); ok {
		fmt.Println(p.Print())
	}
	if s, ok := v.(Scanner); ok {
		fmt.Println(s.Scan())
	}
	if f, ok := v.(Faxer); ok {
		fmt.Println(f.Fax())
	}
}

func main() {
	Check(&AllInOne{Model: "HP_578"})
	Check(&SimplePrinter{Model: "Sony-987"})
}
