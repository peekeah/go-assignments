package main

import "errors"

type Book struct {
	Id              int
	Title           string
	Author          string
	Year            int
	AvailableCopies int
}

type Library []Book

func InitializeLibrary(books []Book) *Library {
	return (*Library)(&books)
}

func (l *Library) AddBook(book Book) {
	*l = append(*l, book)
}

func (l *Library) GetBookById(id int) Book {
	for _, book := range *l {
		if book.Id == id {
			return book
		}
	}
	return Book{}
}

func (l *Library) GetBooks() []Book {
	return *l
}

func (l *Library) DeleteBook(bookId int) Book {
	for id, book := range *l {
		if book.Id == bookId {
			*l = append((*l)[:id], (*l)[id+1:]...)
			return book
		}
	}
	return Book{}
}

func (l *Library) UpdateBook(book Book) Book {
	for id, entry := range *l {
		if entry.Id == book.Id {
			(*l)[id] = book
			return book
		}
	}
	return Book{}
}

func (l *Library) GetAvailableBooks() Library {
	availableBooks := Library{}

	for _, book := range *l {
		if book.AvailableCopies > 0 {
			availableBooks = append(availableBooks, book)
		}
	}

	return availableBooks
}

func (l *Library) BorrowBook(bookId int) (Book, error) {
	for id, book := range *l {
		if book.Id == bookId && book.AvailableCopies == 0 {
			return Book{}, errors.New("Book not available")
		}
		if book.Id == bookId {
			(*l)[id].AvailableCopies--
			return (*l)[id], nil
		}
	}
	return Book{}, errors.New("Book not found")
}

func (l *Library) ReturnBook(bookId int) (Book, error) {
	for id, book := range *l {
		if book.Id == bookId {
			(*l)[id].AvailableCopies++
			return (*l)[id], nil
		}
	}
	return Book{}, errors.New("Book not found")
}
