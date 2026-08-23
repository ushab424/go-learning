package main

import "fmt"

type Subject struct {
	Name  string
	Grade int
}

type GradeBook struct {
	StudentName string
	Subjects    []Subject
}

func (g GradeBook) AddSubject(s Subject) GradeBook {
	g.Subjects = append(g.Subjects, s)
	return g
}

func (g GradeBook) Average() float64 {
	var SummaryGrade int
	MeanValue := 0.0
	for _, Value := range g.Subjects {
		SummaryGrade += Value.Grade
	}
	MeanValue = float64(SummaryGrade) / float64(len(g.Subjects))
	return MeanValue
}

func (g GradeBook) Best() string {
	var BestSubject string
	MaxGrade := 0
	for _, Value := range g.Subjects {
		if Value.Grade > MaxGrade {
			MaxGrade = Value.Grade
			BestSubject = Value.Name
		}
	}
	return BestSubject
}

func (g GradeBook) Worst() string {
	var WorstSubject string
	WorstGrade := 5
	for _, Value := range g.Subjects {
		if Value.Grade < WorstGrade {
			WorstGrade = Value.Grade
			WorstSubject = Value.Name
		}
	}
	return WorstSubject
}

func main() {
	Student := GradeBook{StudentName: "Ivan"}
	Student = Student.AddSubject(Subject{"Bio", 4})
	Student = Student.AddSubject(Subject{"Hist", 2})
	Student = Student.AddSubject(Subject{"Eng", 4})
	Student = Student.AddSubject(Subject{"Math", 5})
	fmt.Println(Student.Average())
	fmt.Println(Student.Best())
	fmt.Println(Student.Worst())
}
