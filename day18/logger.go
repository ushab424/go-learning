package main

import "fmt"

type Logger interface {
	Log(text string)
}

type ConsoleLogger struct{}

type FileLogger struct {
	FileName string
}

func (c *ConsoleLogger) Log(text string) {
	fmt.Println(text)
}

func (f *FileLogger) Log(text string) {
	fmt.Println(f.FileName, text)
}

func Process(l Logger) {
	l.Log("start")
	l.Log("done")
}

func main() {
	file := FileLogger{FileName: "winter"}
	console := &ConsoleLogger{}
	Process(console)
	Process(&file)
}
