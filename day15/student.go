package main

import "fmt"

type Student struct {
	Name   string
	Grades []int
}

func (s Student) AddGrade(grade int) Student {
	s.Grades = append(s.Grades, grade)
	return s
}

func (s Student) Average() float64 {
	var SummaryGrade int
	MeanValue := 0.0
	for _, Value := range s.Grades {
		SummaryGrade += Value
	}
	MeanValue = float64(SummaryGrade) / float64(len(s.Grades))
	return MeanValue
}

func (s Student) IsPassing() bool {
	if s.Average() >= 3.0 {
		return true
	}
	return false
}

func main() {
	student := Student{Name: "Ivan"}
	student = student.AddGrade(5)
	student = student.AddGrade(4)
	student = student.AddGrade(3)
	student = student.AddGrade(5)
	student = student.AddGrade(2)
	fmt.Println(student.Average())
	fmt.Println(student.IsPassing())
}
