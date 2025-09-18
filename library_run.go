package main

import "fmt"

func RunLibrary() {
	initialBooks := []Book{
		{Id: 1, Title: "The Alchemist", Author: "Paulo Coelho", Year: 1997, AvailableCopies: 3},
		{Id: 2, Title: "1984", Author: "George Orwell", Year: 1949, AvailableCopies: 5},
		{Id: 3, Title: "To Kill a Mockingbird", Author: "Harper Lee", Year: 1960, AvailableCopies: 0},
		{Id: 4, Title: "Pride and Prejudice", Author: "Jane Austen", Year: 1813, AvailableCopies: 2},
		{Id: 5, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Year: 1925, AvailableCopies: 6},
		{Id: 6, Title: "Moby Dick", Author: "Herman Melville", Year: 1851, AvailableCopies: 0},
		{Id: 7, Title: "War and Peace", Author: "Leo Tolstoy", Year: 1869, AvailableCopies: 2},
		{Id: 8, Title: "The Catcher in the Rye", Author: "J.D. Salinger", Year: 1951, AvailableCopies: 4},
		{Id: 9, Title: "The Hobbit", Author: "J.R.R. Tolkien", Year: 1937, AvailableCopies: 5},
		{Id: 10, Title: "Brave New World", Author: "Aldous Huxley", Year: 1932, AvailableCopies: 3},
	}

	// Library management
	library := InitializeLibrary(initialBooks)

	// Add book
	library.AddBook(Book{
		Id:              11,
		Title:           "Two States",
		Author:          "Chetan Bhagat",
		Year:            2007,
		AvailableCopies: 9,
	})

	// Update book
	library.UpdateBook(Book{
		Id:              5,
		Title:           "winner stands alone",
		Author:          "Paulo Coelho",
		Year:            2001,
		AvailableCopies: 3,
	})

	// Delete Book
	library.DeleteBook(10)

	// Get GetBookById
	fmt.Println(library.GetBookById(6))

	// Get Get All Books
	fmt.Println(library)

	// Get available Books
	fmt.Println(library.GetAvailableBooks())

	// Borrow book
	fmt.Println(library.BorrowBook(7))
	fmt.Println(library.BorrowBook(7))

	// Return book
	fmt.Println(library.ReturnBook(11))
	fmt.Println(library.ReturnBook(11))

	fmt.Println(library.GetAvailableBooks())
}
