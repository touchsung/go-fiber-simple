// service/book_service.go

package service

import (
	"github.com/touchsung/go-fiber-simple/model"
	"gorm.io/gorm"
)

// BookService provides methods to interact with the Book model.
type BookService struct {
	DB *gorm.DB
}

// NewBookService creates a new instance of BookService.
func NewBookService(db *gorm.DB) *BookService {
	return &BookService{DB: db}
}

// GetAllBooks retrieves all books from the database.
func (s *BookService) GetAllBooks() ([]model.Book, error) {
	var books []model.Book
	err := s.DB.Find(&books).Error
	return books, err
}

// GetBookByID retrieves a specific book by ID from the database.
func (s *BookService) GetBookByID(id uint) (model.Book, error) {
	var book model.Book
	err := s.DB.First(&book, id).Error
	return book, err
}

// CreateBook creates a new book in the database.
func (s *BookService) CreateBook(book *model.Book) error {
	return s.DB.Create(book).Error
}

// UpdateBook updates an existing book in the database.
func (s *BookService) UpdateBook(existingBook *model.Book, updatedBook *model.Book) error {
	return s.DB.Model(existingBook).Updates(updatedBook).Error
}

// DeleteBook deletes a book from the database.
func (s *BookService) DeleteBook(book *model.Book) error {
	return s.DB.Delete(book).Error
}
