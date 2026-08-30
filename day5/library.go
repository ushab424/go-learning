package main

import "fmt"

type Book struct {
	Title string
	Pages int
}

type Library struct {
	Books []Book
}

func (l Library) Add(b Book) Library {
	l.Books = append(l.Books, b)
	return l
}

func (l Library) Count() int {
	return len(l.Books)
}

func (l Library) Biggest() string {
	var BigBook string
	MaxPages := 0
	for _, Name := range l.Books {
		if Name.Pages > MaxPages {
			BigBook = Name.Title
			MaxPages = Name.Pages
		}
	}
	return BigBook
}

func main() {
	library := Library{}
	library = library.Add(Book{Title: "Gatsby", Pages: 250})
	library = library.Add(Book{Title: "Pencil", Pages: 1250})
	library = library.Add(Book{Title: "Christmas", Pages: 2250})
	fmt.Println(library.Count())
	fmt.Println(library.Biggest())
}
