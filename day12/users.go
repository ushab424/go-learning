package main

import "fmt"

func main() {
	users := map[string]string{
		"alice":   "admin",
		"bob":     "user",
		"charlie": "moderator",
		"dave":    "user",
		"eve":     "admin",
	}
	byRole := map[string]([]string){}
	for name, role := range users {
		byRole[role] = append(byRole[role], name)
	}
	for role, name1 := range byRole {
		fmt.Println(role, name1)
	}
}
