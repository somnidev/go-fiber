package services

import (
	"github.com/google/uuid"
	"github.com/somnidev/go-fiber/model"
)

type BookService struct {
	books map[string]model.Book
}

func NewBookService() (*BookService, error) {
	uuid1 := uuid.New().String()
	uuid2 := uuid.New().String()
	uuid3 := uuid.New().String()
	bs := map[string]model.Book{
		uuid1: {ID: uuid1, Title: "Learning Go: An Idiomatic Approach to Real-World Go Programming", Author: "Jon Bodner"},
		uuid2: {ID: uuid2, Title: "Introduction to Algorithms, fourth edition 4th", Author: "Thomas H. Cormen"},
		uuid3: {ID: uuid3, Title: "Clean Code: A Handbook of Agile Software Craftsmanship", Author: "Robert C. Martin"},
	}
	return &BookService{books: bs}, nil
}

func (bookService *BookService) ListBooks() []model.Book {
	books := make([]model.Book, 0, len(bookService.books))
	for _, value := range bookService.books {
		books = append(books, value)
	}
	return books
}

func (bookService *BookService) CreateBook(book model.Book) model.Book {
	uuid := uuid.New().String()
	book.ID = uuid
	bookService.books[uuid] = book
	return book
}
