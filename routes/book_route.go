// route/book_route.go

package route

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/touchsung/go-fiber-simple/model"
	"github.com/touchsung/go-fiber-simple/service"
)

// SetupBookRoutes sets up routes related to the Book model.
func SetupBookRoutes(app *fiber.App, bookService *service.BookService) {
	// GET all books
	app.Get("/books", func(c *fiber.Ctx) error {
		books, err := bookService.GetAllBooks()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve books"})
		}
		return c.JSON(books)
	})

	// GET a specific book by ID
	app.Get("/books/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
		}
		book, err := bookService.GetBookByID(uint(id))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		}
		return c.JSON(book)
	})

	// Create a new book
	app.Post("/books", func(c *fiber.Ctx) error {
		var book model.Book
		if err := c.BodyParser(&book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		if err := bookService.CreateBook(&book); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create book"})
		}
		return c.Status(fiber.StatusCreated).JSON(book)
	})

	// Update a book
	app.Put("/books/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
		}
		var existingBook model.Book
		result := bookService.DB.First(&existingBook, id)
		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		}
		var updatedBook model.Book
		if err := c.BodyParser(&updatedBook); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		}
		if err := bookService.UpdateBook(&existingBook, &updatedBook); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update book"})
		}
		return c.JSON(existingBook)
	})

	// Delete a book
	app.Delete("/books/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
		}
		var book model.Book
		book.ID = uint(id)
		if err := bookService.DeleteBook(&book); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete book"})
		}
		return c.Status(fiber.StatusNoContent).Send(nil)
	})
}
