package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/touchsung/go-fiber-simple/handler"
	"github.com/touchsung/go-fiber-simple/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Starting server...")

	// Initialize the database connection
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5434 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// AutoMigrate the schema
	err = db.AutoMigrate(&model.Book{})
	if err != nil {
		panic("failed to migrate database")
	}

	// Create a new Fiber instance
	app := fiber.New()

	// Initialize the book handler
	bookHandler := handler.NewBookHandler(db)

	// Setup routes
	bookHandler.SetupRoutes(app)

	// Start the server
	err = app.Listen(":4000")
	if err != nil {
		panic("failed to start server")
	}

	// Close the database connection when the application exits
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get DB connection")
	}
	defer sqlDB.Close()
}
