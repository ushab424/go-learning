package main

import "fmt"

type Task struct {
	Title string
	Done  bool
}

type ToDooList struct {
	Tasks []Task
}

func (t ToDooList) Add(title string) ToDooList {
	t.Tasks = append(t.Tasks, Task{Title: title, Done: false})
	return t
}

func (t ToDooList) Complete(title string) ToDooList {
	for i := range t.Tasks {
		if t.Tasks[i].Title == title {
			t.Tasks[i].Done = true
		}
	}
	return t
}

func (t ToDooList) Pending() ToDooList {
	NotComplited := []Task{}
	for i := range t.Tasks {
		if t.Tasks[i].Done == false {
			NotComplited = append(NotComplited, t.Tasks[i])
		}
	}
	return ToDooList{Tasks: NotComplited}
}

func main() {
	TaskList := ToDooList{}
	TaskList = TaskList.Add("Постирать")
	TaskList = TaskList.Add("Помыть посуду")
	TaskList = TaskList.Add("Продукты")
	TaskList = TaskList.Add("Учеба")

	TaskList = TaskList.Complete("Постирать")
	TaskList = TaskList.Complete("Учеба")

	pending := TaskList.Pending()
	for _, task := range pending.Tasks {
		fmt.Println(task.Title)
	}
}
