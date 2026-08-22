package main

import "fmt"

type Employee struct {
	Name     string
	Position string
	Salary   float64
}

func (E Employee) Raise(Percent float64) Employee {
	E.Salary += (E.Salary * Percent)
	return E
}

func (E Employee) Info() {
	fmt.Println(E.Name, E.Position, E.Salary)
}

func main() {
	Person := Employee{Name: "Ivan Dubravin", Position: "Manager", Salary: 95000}
	Person = Person.Raise(0.15)
	Person.Info()
}
