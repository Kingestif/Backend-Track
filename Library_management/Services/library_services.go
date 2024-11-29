package Services

import (
	"errors"
	"Library_management/Models"
)

type LibraryManager interface{
	AddBook(book Models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []Models.Book
	ListBorrowedBooks(memberID int) []Models.Book
}


//Library struct to hold books and members 
type Library struct {
	Books map[int] *Models.Book			//(bookid: title, author, status)
	Members  map[int] *Models.Member		//(memberid: name, borrowedbooks)
}

//Just instantiating the Library struct with empty maps(only called once), since go doesn't have constructors we use this method
func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]*Models.Book),
		Members: make(map[int]*Models.Member),
	}
}


// passing l *Library makes Library available for the function
func(l *Library) AddBook(book Models.Book) {
	l.Books[book.Id] = &book
}

func(l *Library) RemoveBook(bookID int){
	delete(l.Books, bookID)
}

func(l *Library) BorrowBook(bookID int, memberID int) error{
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}

	member, exists := l.Members[memberID]
	if !exists {
		member = &Models.Member{Id: memberID}
	}

	// check if book is available(status)
	if book.Status == "Borrowed"{
		return errors.New("book not available")
	}

	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, *book)

	// finally update the Library
	l.Books[bookID] = book
    l.Members[memberID] = member

	return nil
} 

func(l *Library) ReturnBook(bookID int, memberID int) error{
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not exist")
	}

	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not exist")
	}

	if book.Status == "Available" {
		return errors.New("book is already available")
	}

	
	for i,bk := range member.BorrowedBooks{		//its array not map
		if bk.Id == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
		}
	}
	
	book.Status = "Available"

	l.Books[bookID] = book
    l.Members[memberID] = member

	return nil
} 

func(l *Library) ListAvailableBooks() []Models.Book{
	var AvailableBooks []Models.Book
	for _,bk := range l.Books{
		if bk.Status == "Available"{
			AvailableBooks = append(AvailableBooks, *bk)
		}
	}

	return AvailableBooks
}

func(l *Library) ListBorrowedBooks(memberID int) []Models.Book{
	member, exists := l.Members[memberID]

	if !exists {
		return nil
	}

	return member.BorrowedBooks

}