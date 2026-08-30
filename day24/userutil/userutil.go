package userutil

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) (*User, error) {
	if name == "" || age < 0 {
		return nil, errors.New("incorrect data")
	} else {
		return &User{Name: name, Age: age}, nil
	}
}
func FormatUser(u *User) string {
	return fmt.Sprintf("Name: %s, Age: %d", u.Name, u.Age)
}
