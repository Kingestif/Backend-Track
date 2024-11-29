//Controller layer recives input and connect to service Layer

package Controllers

import (
	"fmt"
	"Library_management/Models"
	"Library_management/Services"
)


type LibraryController struct {
	service *Services.Library
}

func NewLibraryController() *LibraryController {
	return &LibraryController{
		service: Services.NewLibrary(),
	}
}

func(lc *LibraryController) AddBook(){
	var id int
	var author, title string

	fmt.Println("Insert Book ID?")
	fmt.Scanln(&id)

	fmt.Println("Insert Book Author?")
	fmt.Scanln(&author)

	fmt.Println("Insert Book Title?")
	fmt.Scanln(&title)


	book := Models.Book{Author: author, Title: title, Id: id, Status: "Available"}
	lc.service.AddBook(book)
}

func(lc *LibraryController) RemoveBook(){
	var id int

	fmt.Println("Insert Book ID to be removed?")
	fmt.Scanln(&id)

	lc.service.RemoveBook(id)
}

func(lc *LibraryController) BorrowBook(){
	var bookid int
	var memberid int

	fmt.Println("What is your ID")
	fmt.Scanln(&memberid)

	fmt.Println("What is the book ID")
	fmt.Scanln(&bookid)

	err := lc.service.BorrowBook(bookid, memberid)

	if err != nil {
		fmt.Println("Error borrowing book:", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func(lc *LibraryController) ReturnBook(){
	var bookid int
	var memberid int

	fmt.Println("What is your ID")
	fmt.Scanln(&memberid)

	fmt.Println("What is the book ID")
	fmt.Scanln(&bookid)

	err := lc.service.ReturnBook(bookid,memberid)
	
	if err != nil {
		fmt.Println("Error returning book:", err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func(lc *LibraryController) ListAvailableBooks(){
	books := lc.service.ListAvailableBooks()
	if len(books) == 0{
		fmt.Println("No available book")
	}else{
		for _,book := range books{
			fmt.Println("ID", book.Id, "Title", book.Title, "Author", book.Author)
		} 
	}

}

func(lc *LibraryController) ListBorrowedBooks(){
	var memberid int

	fmt.Println("What is your ID")
	fmt.Scanln(&memberid)

	books := lc.service.ListBorrowedBooks(memberid)

	if len(books) == 0{
		fmt.Println("No Borrowed Books")
	}else{
		for _,book := range books{
			fmt.Println("ID", book.Id, "Title", book.Title, "Author", book.Author)
		} 
	}
}