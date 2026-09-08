package main

import (
	"fmt"
	"reflect"
)

func printInfo(x any) {
	resultval := reflect.ValueOf(x)
	resulttype := reflect.TypeOf(x)
	fmt.Println("type:", resulttype, "| value:", resultval)
}

func main() {
	printInfo(3)
	printInfo(3.33)
	printInfo("abc")
}
