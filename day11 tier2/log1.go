package main

import "log"

func main() {
	log.Println("server started")
	// как Println, но автоматически добавляет время и дату
	log.Printf("port: %d", 8080)

	log.SetPrefix("[APP] ")
	// добавляет префикс перед каждым сообщением.
	log.Println("with prefix")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// настраивает формат. Ldate — дата, Ltime — время, Lshortfile — имя файла и строка кода.
	log.Println("with file info")
}

/*
Ещё два важных метода:
log.Fatal("msg") — печатает сообщение и завершает программу с кодом 1. Используй когда ошибка критическая и продолжать нельзя.
log.Panic("msg") — печатает сообщение и вызывает panic. Можно перехватить через recover.
*/
