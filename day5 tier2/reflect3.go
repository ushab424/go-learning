package main

import "reflect"

type User1 struct {
	Name  string
	Age   int
	Email string
}

func SetField(x any, fieldName string, newValue any) {
	v := reflect.ValueOf(x)
	v.Elem().FieldByName(fieldName).Set(reflect.ValueOf(newValue))
}

func main() {
	user1 := User1{Name: "alex", Age: 25, Email: "slkveij@.fo"}
	SetField(&user1, "Name", "bob")
}
