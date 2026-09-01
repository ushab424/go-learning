package storage

import (
	"bufio"
	"fmt"
	"os"
	"phonebook/contact"
	"strconv"
	"strings"
)

func SaveContacts(filename string, contacts []*contact.Contact) error { // Сохранение контакта
	people, err := os.Create(filename)
	if err != nil { // проверка на ошибку создания файла
		fmt.Println("Ошибка создания файла:", err)
		return err
	}
	defer people.Close()            // обязательно закрыть файл
	for _, cont := range contacts { // записываем в созданный файл контакт
		people.WriteString(fmt.Sprintf("%d|%s|%s|%s|%s\n", cont.ID, cont.Name, cont.Phone, cont.Email, cont.Group))
	}
	return nil
}

func LoadContacts(filename string) ([]*contact.Contact, error) { // загрузка контактов
	file, err := os.Open(filename) // открывает файл
	if err != nil {                // проверяем файл на ошибку открытия
		fmt.Println("file don't open:", err)
		return nil, err
	}
	defer file.Close()                // обязательно закрываем
	scanner := bufio.NewScanner(file) // сканируем файл построчно
	var contacts []*contact.Contact   // создаем слайс куда будем записыва контакты
	for scanner.Scan() {              // приступаем к сканированию
		line := scanner.Text()            // читаем построчно
		parts := strings.Split(line, "|") // разделяем строки "|"
		id, err := strconv.Atoi(parts[0]) // конвертируем id из строки в int
		if err != nil {                   // проверка конвертанции на ошибку
			fmt.Println("error:", err)
		}
		if err := scanner.Err(); err != nil { // проверка на ошибку сканера
			return nil, err
		}
		c := contact.AddContact(id, parts[1], parts[2], parts[3], parts[4]) // используем функцию добавления контакта
		contacts = append(contacts, c)                                      // добавляем контакт в файл
	}
	return contacts, nil
}
