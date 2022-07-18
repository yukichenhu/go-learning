package entity

type Book struct {
	Title   string
	Author  string
	Subject string
	BookId  int
}

type Author struct {
	Name  string
	Age   int
	Sex   string
	Books []Book
}
