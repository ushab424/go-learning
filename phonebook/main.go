package main

import (
	"phonebook/contact"
	"phonebook/menu"
	"phonebook/storage"
)

func main() {
	filename := "ContactList.txt"                   // присваиваем имя файлу с контактами
	contacts, err := storage.LoadContacts(filename) // активируем функцию загрузки контактов
	if err != nil {
		contacts = []*contact.Contact{}
	}
	menu.ShowMenu(&contacts, filename)       // активируем наше меню
	storage.SaveContacts(filename, contacts) // при выходе сохраняем изменения
}
