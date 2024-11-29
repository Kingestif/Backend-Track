package main

import (
	"fmt"
	"Library_management/Controllers"
)

func main() {
	controller := Controllers.NewLibraryController()

	for {
		fmt.Println("Welcome To Library")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Remove an existing book")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. List all available books")
		fmt.Println("6. List all borrowed books")
		fmt.Println("7. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			controller.AddBook()
		case 2:
			controller.RemoveBook()
		case 3:
			controller.BorrowBook()
		case 4:
			controller.ReturnBook()
		case 5:
			controller.ListAvailableBooks()
		case 6:
			controller.ListBorrowedBooks()
		case 7:
			return
		default:
			fmt.Println("Invalid choice please choose again")
		}
	}
}
