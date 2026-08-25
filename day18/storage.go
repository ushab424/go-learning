package main

import "fmt"

type Storage interface {
	Save(key string, value string)
	Get(key string) string
}

type MemoryStorage struct {
	Data map[string]string
}

type FileStorage struct{}

func (m *MemoryStorage) Save(key string, value string) {
	m.Data[key] = value
}

func (m *MemoryStorage) Get(key string) string {
	return m.Data[key]
}

func (f *FileStorage) Save(key string, value string) {
	fmt.Println("saved to file:", key)
}

func (f *FileStorage) Get(key string) string {
	return "file" + key
}

func Demo(s Storage) {
	s.Save("name", "ivan")
	fmt.Println(s.Get("name"))
}

func main() {
	mem := &MemoryStorage{Data: make(map[string]string)}
	file := &FileStorage{}
	Demo(mem)
	Demo(file)
}
