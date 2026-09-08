package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func printFields(x any) {
	v := reflect.ValueOf(x)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fmt.Println(t.Field(i).Name, "=", v.Field(i))
	}
}

func main() {
	u1 := User{Name: "alex", Age: 25, Email: "slkveij@.fo"}
	printFields(u1)
}
