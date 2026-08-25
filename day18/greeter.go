package main

import "fmt"

type Greeter interface {
	Greet(name string) string
}

type Formal struct{}
type Casual struct{}
type Robot struct {
	ID string
}

func (f *Formal) Greet(name string) string {
	return fmt.Sprintf("Good day, %s", name)
}
func (c *Casual) Greet(name string) string {
	return fmt.Sprintf("Hey, %s", name)
}
func (r *Robot) Greet(name string) string {
	return fmt.Sprintf("Unit %s greets, %s", r.ID, name)
}
func welcome(g Greeter, name string) {
	fmt.Println(g.Greet(name))
}
func main() {
	greet := Formal{}
	greet1 := Casual{}
	greet2 := Robot{ID: "xd-3000"}
	welcome(&greet, "Ivan")
	welcome(&greet1, "Ivan")
	welcome(&greet2, "Ivan")
}
