package menu

import (
	"bufio"
	"fmt"
	"os"
	"phonebook/contact"
	"phonebook/validator"
	"sort"
	"strconv"
	"strings"
)

func ShowMenu(contacts *[]*contact.Contact, filename string) { // функция использования меню и всех функций приложения
	scanner := bufio.NewScanner(os.Stdin) // активируем сканер
	for {                                 // бесконечный цикл
		fmt.Println("1-Добавить 2-удалить 3-показать все 4-поиск 5-редактировать контакт 6-фильтрация по группе 7-статистика по группам 8-эксопорт в .csv 9-сортировка 10-выход") // варианты использования
		scanner.Scan()                                                                                                                                                            // сканируем ответ
		choice := scanner.Text()                                                                                                                                                  // читаем ответ и используем функции
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
			found := false                // создаем флаг
			for i, r := range *contacts { // крутим список контактов
				if r.ID == idNum { // сверяем сходство
					found = true                                              // меняем значение флага
					*contacts = append((*contacts)[:i], (*contacts)[i+1:]...) // склеиваем строку без этого контакта
					fmt.Println("Contact delete")
					break // выходим из цикла
				}
			}
			if !found { // если контакт не найден
				fmt.Println("Contact not found.") // печатаем ошибку
			}
		case "3": // вывод всех контактов
			fmt.Println("Показать все")
			if len(*contacts) == 0 { // если контакты отстутсвуют
				fmt.Println("0 contacts")
			} else {
				for _, r := range *contacts { // крутим список
					fmt.Println(r) // и форматированно выводим всех
				}
			}
		case "4": // поиск по строке
			fmt.Println("Enter string:")
			scanner.Scan()
			searchstring := scanner.Text() // присваиваем введеное слова переменной
			for _, r := range *contacts {  // крутим контакты
				if strings.Contains(r.Name, searchstring) || strings.Contains(r.Phone, searchstring) || strings.Contains(r.Email, searchstring) || strings.Contains(r.Group, searchstring) {
					fmt.Println(r) // если хоть одна строка свопадает, форматированно выводим
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
			found := false                // создаем флаг
			for _, r := range *contacts { // крутим контакты
				if r.ID == id { // поиск совпадения ID
					found = true                                       // меняем значение флага
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
			if !found { //если ID не найден
				fmt.Println("ID not found!")
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
		case "7": // функция статистики по группам
			stat := make(map[string]int)   // создаем мапу где каждый ключ это группа
			for _, gr := range *contacts { // крутим контакты и добавляем в мапу
				stat[gr.Group]++
			}
			fmt.Printf("All contacts: %d\n", len(*contacts)) // общее количество контактов
			for i, val := range stat {                       // крутим мапу
				fmt.Printf("Group: %s. | Quat: %d", i, val) // выводим группу и количество контактов
			}
		case "8":
			file, err := os.Create("ContactList.csv") //создаем файл .csv
			if err != nil {                           // проверяем на ошибку создания файла
				fmt.Println("error:", err)
			}
			defer file.Close()                             // закрываем файл обязательно!!!
			fmt.Fprintf(file, "ID,Name,Phone,Email,Group") // пишем заголовки в файле
			for _, r := range *contacts {                  // крутим контакты
				fmt.Fprintf(file, "%d,%s,%s,%s,%s\n", r.ID, r.Name, r.Phone, r.Email, r.Group) // печатаем данные в каждой строке по каждому контакту
			}
			fmt.Println("Export done!") // выводим уведомление о завершении экспорта
		case "9": // функция сортировки по имени или ID
			fmt.Println("Select a sorting option: 1 – by name, 2 – by ID:") // запрашиваем вид сортировки
			scanner.Scan()
			variat := scanner.Text()
			switch variat {
			case "1": // сортировка по имени
				sort.Slice(*contacts, func(i, j int) bool { // сортируем слайс и возвращаем true/false
					return (*contacts)[i].Name < (*contacts)[j].Name
				})
			case "2": // сортировка по ID
				sort.Slice(*contacts, func(i, j int) bool { // сортируем слайс и возвращаем true/false
					return (*contacts)[i].ID < (*contacts)[j].ID
				})
			}
			for _, r := range *contacts { // крутим отсортированный слайс и возвращаем его
				fmt.Println(r)
			}
		case "10":
			fmt.Println("Выход")
			return // выход из бесконечного цикла меню
		}
		if err := scanner.Err(); err != nil { // проверка на ошибку сканера ответа меню
			return
		}
	}
}
