package contact

type Contact struct { // Создаем структуру контакта
	ID    int
	Name  string
	Phone string
	Email string
	Group string
}

func AddContact(id int, name, phone, email, group string) *Contact {
	return &Contact{ID: id, Name: name, Phone: phone, Email: email, Group: group}
}

// Функция добавления контакта
