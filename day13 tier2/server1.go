package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Хендлер - это функция которая отвечает на запрос.
	// Каждый хендлер принимает два аргумента: w и r!

	// r *http.Request — это конверт с запросом от клиента. Из него ты достаёшь информацию.

	// Все что мы пишем в w, клиент увидит в браузере.
	fmt.Fprintf(w, "hello world!")
	// пишем текст в w и выводим его в браузер. (или просто отдаем данные)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(time.Now().Format("15:04:05"))
	// json.NewEncoder(w).Encode(данные) — пишет JSON.
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// r.Method — какой метод (GET, POST).
	msg := r.URL.Query().Get("msg")
	// r.URL.Query().Get("name") — параметр из URL после знака ?.

	// r.Body — тело запроса (данные которые клиент отправил в POST).
	if msg == "" {
		msg = "empty"
	}
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set(...) — устанавливает заголовок (тип контента)

	// // w.WriteHeader(код) — устанавливает статус-код (200, 201, 400).
	json.NewEncoder(w).Encode(map[string]string{"message": msg})
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	// http.HandleFunc("/hello", helloHandler) — говорит серверу: "когда кто-то зайдёт на /hello, вызови функцию helloHandler". Это как таблица: путь → функция.
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/echo", echoHandler)

	fmt.Println("server started at :8080")
	http.ListenAndServe(":8080", nil)
	// http.ListenAndServe(":8080", nil) — запускает сервер. Он начинает слушать порт 8080 и ждёт запросы. Программа зависает здесь навсегда — пока не нажмёшь Ctrl+C.
}
