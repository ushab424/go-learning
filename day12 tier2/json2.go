package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price,omitempty"`
	// omitempty — если значение пустое (0 для int, "" для string), поле не попадёт в JSON. У p2 Price=0, поэтому в JSON его не будет.
	Secret string `json:"-"`
	// json:"-" — поле никогда не попадёт в JSON. Secret не появится в выводе. Используют для паролей, токенов, внутренних данных.

}

func main() {
	p1 := Product{ID: 1, Name: "Phone", Price: 999, Secret: "hidden"}
	p2 := Product{ID: 2, Name: "Case", Price: 0, Secret: "hidden"}

	data1, _ := json.Marshal(p1)
	data2, _ := json.Marshal(p2)

	fmt.Println(string(data1))
	fmt.Println(string(data2))
}
