package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wokacz/hermod/cmd/users/handler"
	"github.com/wokacz/hermod/model"
	"github.com/wokacz/hermod/pkg/acl"
	"github.com/wokacz/hermod/pkg/database"
	"github.com/wokacz/hermod/pkg/env"
	"github.com/wokacz/hermod/pkg/jwt"
	"gorm.io/gorm"
)

// setupRoutes function is used to setup the routes for the application.
func setupRoutes(app *fiber.App) {
	// Setup the health check route.
	app.Get("/__health", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
	// Setup the JWT middleware.
	secret := env.Get("JWT_SECRET", "top-secret")
	app.Use(jwt.New(jwt.Config{
		Secret: secret,
	}))
	// Setup the ACL middleware.
	app.Use(acl.New(acl.Config{
		Roles: []string{"admin"},
	}))
	// Setup the routes.
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
		panic(err.Error())
	}
	// Migrate database.
	err = database.DB.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.AutoMigrate(&model.User{})
		return err
	})
	// If the database migration fails, panic.
	if err != nil {
		panic(err.Error())
	}
}

// main function is used to start the application.
func main() {
	// Create a new Fiber application.
	app := fiber.New()
	// Setup the routes.
	setupRoutes(app)
	// Listen on port 3005.
	err := app.Listen(":3010")
	// If the application fails to listen, panic.
	if err != nil {
		panic(err)
	}
}
