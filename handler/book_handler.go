// handler/book_handler.go

package handler

import (
	"github.com/gofiber/fiber/v2"
	route "github.com/touchsung/go-fiber-simple/routes"
	"github.com/touchsung/go-fiber-simple/service"
	"gorm.io/gorm"
)

// BookHandler handles HTTP requests related to the Book model.
type BookHandler struct {
	BookService *service.BookService
}

// NewBookHandler creates a new instance of BookHandler.
func NewBookHandler(db *gorm.DB) *BookHandler {
	return &BookHandler{
		BookService: service.NewBookService(db),
	}
}

// SetupRoutes sets up routes for BookHandler.
func (h *BookHandler) SetupRoutes(app *fiber.App) {
	route.SetupBookRoutes(app, h.BookService)
}
