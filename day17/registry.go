package main

import (
	"fmt"
)

type User struct {
	Name   string
	Email  string
	Active bool
}

type Registry struct {
	Users []User
}

func (r *Registry) Register(Name string, Email string) {
	r.Users = append(r.Users, User{Name: Name, Email: Email, Active: true})
}

func (r *Registry) Deactivate(Email string) {
	for i, user := range r.Users {
		if user.Email == Email {
			r.Users[i].Active = false
		}
	}
}

func (r *Registry) ActiveUser() []User {
	activeuser := []User{}
	for i, user := range r.Users {
		if user.Active == true {
			activeuser = append(activeuser, r.Users[i])
		}
	}
	return activeuser
}

func (r *Registry) FindByEmail(Email string) string {
	for _, user := range r.Users {
		if user.Email == Email {
			return user.Name
		}
	}
	return "Not Found"
}

func main() {
	users := Registry{}
	users.Register("Ivan", "sojfo410@gmail.com")
	users.Register("Olga", "flogoo87@gmail.com")
	users.Register("Julia", "psfjfw78@gmail.com")
	users.Register("Oleg", "pakf9we87@gmail.com")
	users.Deactivate("sojfo410@gmail.com")
	fmt.Println(users.ActiveUser())
	fmt.Println(users.FindByEmail("flogoo87@gmail.com"))
}
