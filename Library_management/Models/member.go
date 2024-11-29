package Models

type Member struct{
	Id int
	Name string
	BorrowedBooks []Book
}