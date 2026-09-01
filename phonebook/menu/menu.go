package menu

import (
	"bufio"
	"fmt"
	"os"
	"phonebook/contact"
	"phonebook/validator"
	"strconv"
	"strings"
)

func ShowMenu(contacts *[]*contact.Contact, filename string) { // функция использования меню и всех функций приложения
	scanner := bufio.NewScanner(os.Stdin) // активируем сканер
	for {                                 // бесконечный цикл
		fmt.Println("1-Добавить 2-удалить 3-показать все 4-поиск 5-редактировать контакт 6-фильтрация по группе 7-выход") // варианты использования
		scanner.Scan()                                                                                                    // сканируем ответ
		choice := scanner.Text()                                                                                          // читаем ответ и используем функции
		switch choice {
		case "1": // добавление контакта
			fmt.Println("Name:") // поочереди просим ввести данные и сканируем их
			scanner.Scan()
			name := scanner.Text()

			fmt.Println("Phone:")
			scanner.Scan()
			phone := scanner.Text()
			if !validator.IsValidPhone(phone) { // пользуемся функцией валидации номера телефона
				fmt.Println("error: invalid phone number")
				continue // переходим к следующим данным но не выходим из цикла
			}

			fmt.Println("Email:")
			scanner.Scan()
			email := scanner.Text()
			if !validator.IsValidEmail(email) { // пользуемся функицей валидации имейла
				fmt.Println("error: invalid email")
				continue // переходим к следующим данным но не выходим из цикла
			}

			fmt.Println("Group:")
			scanner.Scan()
			group := scanner.Text()

			id := len(*contacts) + 1                               // id присваиваем самостоятельно
			c := contact.AddContact(id, name, phone, email, group) // пользуемся функцией добавления контакта
			*contacts = append(*contacts, c)                       // добавляем в структуру
			fmt.Println("Contact save")
		case "2": // удаление контакта
			fmt.Println("Enter ID:")
			scanner.Scan()
			id := scanner.Text()
			idNum, err := strconv.Atoi(id) // переводим id в int
			if err != nil {                // проверка на ошибку перевода в int
				fmt.Println("error:", err)
			}
			for i, r := range *contacts { // крутим список контактов
				if r.ID == idNum { // сверяем сходство
					*contacts = append((*contacts)[:i], (*contacts)[i+1:]...) // склеиваем строку без этого контакта
					fmt.Println("Contact delete")
					break // выходим из цикла
				}
			}
		case "3": // вывод всех контактов
			fmt.Println("Показать все")
			for _, cont := range *contacts { // крутим список
				fmt.Printf("ID: %d | Name: %s | PhoneNum: %s | Email: %s | Group: %s\n", cont.ID, cont.Name, cont.Phone, cont.Email, cont.Group) // и форматированно выводим всех
			}
		case "4": // поиск по строке
			fmt.Println("Enter string:")
			scanner.Scan()
			searchstring := scanner.Text() // присваиваем введеное слова переменной
			for _, r := range *contacts {  // крутим контакты
				if strings.Contains(r.Name, searchstring) || strings.Contains(r.Phone, searchstring) || strings.Contains(r.Email, searchstring) || strings.Contains(r.Group, searchstring) {
					fmt.Printf("ID: %d | Name: %s | Phone: %s | Email: %s | Group: %s\n", r.ID, r.Name, r.Phone, r.Email, r.Group) // если хоть одна строка свопадает, форматированно выводим
				}
			}
		case "5": // редактирование контакта
			fmt.Println("Enter ID:") // запрашиваем ID
			scanner.Scan()
			idContact := scanner.Text()
			id, err := strconv.Atoi(idContact) // переводим строку в int
			if err != nil {                    // проверка перевода на ошибку
				fmt.Println(err)
			}
			for _, r := range *contacts { // крутим контакты
				if r.ID == id { // поиск совпадения ID
					fmt.Println("Select the line you want to change:") // запрашиваем строку для изменения
					scanner.Scan()
					line := scanner.Text()
					switch line { // выбираем строку для изменения
					case "Name":
						fmt.Println("Enter data:") // запрашиваем данные
						scanner.Scan()
						text := scanner.Text()
						r.Name = text // меняем данные
					case "Phone":
						fmt.Println("Enter data:") // запрашиваем данные
						scanner.Scan()
						num := scanner.Text()
						if !validator.IsValidPhone(num) { // пользуемся функцией валидации номера телефона
							fmt.Println("error: invalid phone number")
						} else {
							r.Phone = num // меняем данные
						}
					case "Email":
						fmt.Println("Enter data:") // запрашиваем данные
						scanner.Scan()
						email := scanner.Text()
						if !validator.IsValidEmail(email) { // пользуемся функицей валидации имейла
							fmt.Println("error: invalid email")
						} else {
							r.Email = email // меняем данные
						}
					case "Group":
						fmt.Println("Enter data:") // запрашиваем данные
						scanner.Scan()
						group := scanner.Text()
						r.Group = group // меняем данные
					}
				}
			}
		case "6": // фильтрация контактов по группе
			fmt.Println("Select group:") // запрашиваем группу
			scanner.Scan()
			group := scanner.Text()
			for _, r := range *contacts { // крутим контакты
				if r.Group == group { // если совпадение
					fmt.Println(r.ID, r.Name, r.Phone, r.Email, r.Group) // печатаем контакт
				}
			}
		case "7":
			fmt.Println("Выход")
			return // выход из бесконечного цикла меню
		}
		if err := scanner.Err(); err != nil { // проверка на ошибку сканера ответа меню
			return
		}
	}
}

/*
План на завтра:
1.Статистика — кейс "8": сколько всего контактов, сколько в каждой группе
2.Экспорт в CSV(или JSON, спросить у гпт) — кейс "9": сохранить контакты в формате CSV
3.Сортировка — кейс "10": показать контакты отсортированные по имени или по ID
4.Добавить интерфейс Stringer для Contact — метод String() string
5.Обработка ошибок везде — если файл не найден, если ID не существует при удалении,
если слайс пустой при показе
*/
