package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wokacz/hermod/cmd/roles/handler"
	"github.com/wokacz/hermod/model"
	"github.com/wokacz/hermod/pkg/database"
	"gorm.io/gorm"
)

// main function is used to start the application.
func setupRoutes(app *fiber.App) {
	// Setup the health check route.
	app.Get("/__health", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
	// Create a new group of routes.
	app.Get("/", handler.TakeMany)
	app.Get(":id", handler.TakeOne)
	app.Post("/", handler.Create)
	app.Put("/:id", handler.Update)
	app.Delete("/:id", handler.Delete)
}

// init function is used to initialize the application.
func init() {
	// Initialize database connection.
	err := database.Init()
	// If the database initialization fails, panic.
	if err != nil {
		panic(err)
	}
	// Migrate database.
	err = database.DB.Transaction(func(tx *gorm.DB) (err error) {
		tx.AutoMigrate(&model.Privilege{}, &model.Role{})
		return err
	})
	// If the database migration fails, panic.
	if err != nil {
		panic(err)
	}
}

// main function is used to start the application.
func main() {
	// Create a new Fiber application.
	app := fiber.New()
	// Setup the routes.
	setupRoutes(app)
	// Start the application.
	err := app.Listen(":3005")
	// If the application fails to start, panic.
	if err != nil {
		log.Panic(err)
	}
}
