package validatorr

import "strings"

func IsEmpty(s string) bool {
	if s == "" {
		return true
	}
	return false
}
func IsEmail(s string) bool {
	if strings.Contains(s, "@") && strings.Contains(s, ".") {
		return true
	}
	return false
}
