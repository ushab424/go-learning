package main

import "fmt"

type Task struct {
	Title string
	Done  bool
}

type TooDooList struct {
	Tasks []Task
}

func (t *TooDooList) Add(title string) {
	t.Tasks = append(t.Tasks, Task{Title: title, Done: false})
}

func (t *TooDooList) Complete(title string) {
	for i := range t.Tasks {
		if t.Tasks[i].Title == title {
			t.Tasks[i].Done = true
		}
	}
}

func (t *TooDooList) Pending() []Task {
	NotComplieted := []Task{}
	for i := range t.Tasks {
		if t.Tasks[i].Done == false {
			NotComplieted = append(NotComplieted, t.Tasks[i])
		}
	}
	return NotComplieted
}

func main() {
	todolist := TooDooList{}
	todolist.Add("wash")
	todolist.Add("dry")
	todolist.Add("Learning")
	todolist.Add("cook")
	todolist.Complete("dry")
	todolist.Complete("wash")
	fmt.Println(todolist.Pending())

}
