package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{"name":"Ivan","age":15,"active":true,"scores":[90,85,77]}`

	var data map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &data)

	fmt.Println(data["name"])
	fmt.Println(data["age"])
	fmt.Println(data["active"])
	fmt.Println(data["scores"])
}
