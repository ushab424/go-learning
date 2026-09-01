package validator

import (
	"strings"
)

func IsValidEmail(email string) bool { // проверка email на валидность
	ValidEmail := strings.Contains(email, "@") // на "@"
	ValidPoint := strings.Contains(email, ".") // на "."
	if ValidEmail && ValidPoint {
		return true
	}
	return false
}
func IsValidPhone(phone string) bool { // проверка телефона на валидность
	nums := []rune(phone)                 // переводим строку в руны
	if len(nums) < 10 && len(nums) > 12 { // если оба верны возвращаем false
		return false
	}
	for _, r := range nums { // крутим руны
		if r < '0' || r > '9' { // если цифра не попадает в диапазон возвращаем false
			return false
		}
	}
	return true // если обе верны возвращаем true
}
