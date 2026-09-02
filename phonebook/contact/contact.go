package contact

import "fmt"

type Contact struct { // Создаем структуру контакта
	ID    int
	Name  string
	Phone string
	Email string
	Group string
}

func (c *Contact) String() string { // метод вывода
	return fmt.Sprintf("ID: %d | Name: %s | Phone: %s | Email: %s | Group: %s", c.ID, c.Name, c.Phone, c.Email, c.Group)
}

func AddContact(id int, name, phone, email, group string) *Contact {
	return &Contact{ID: id, Name: name, Phone: phone, Email: email, Group: group}
}

// Функция добавления контакта
