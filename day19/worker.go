package main

import "fmt"

type Worker interface {
	Work() string
}
type Rester interface {
	Rest() string
}

type Human struct {
	Name string
}

type Robot struct {
	Model string
}

func (h *Human) Work() string {
	return h.Name + " is working"
}
func (h *Human) Rest() string {
	return h.Name + " is chill"
}

func (r *Robot) Work() string {
	return r.Model + " is working"
}

func Describe(w Worker) {
	switch value := w.(type) {
	case *Human:
		fmt.Println(value.Work())
		fmt.Println(value.Rest())
	case *Robot:
		fmt.Println(value.Work())
	default:
		fmt.Println("NotFound")
	}
}

func main() {
	Describe(&Human{Name: "Ivan"})
	Describe(&Robot{Model: "xd3000"})
}
