package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wokacz/hermod/cmd/boards/handler"
	"github.com/wokacz/hermod/model"
	"github.com/wokacz/hermod/pkg/database"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/__health", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	app.Get("/", handler.TakeMany)
	app.Get(":id", handler.TakeOne)
	app.Post("/", handler.Create)
	app.Put("/:id", handler.Update)
	app.Delete("/:id", handler.Delete)
}

func init() {
	// Initialize the database.
	err := database.Init()
	if err != nil {
		panic(err.Error())
	}
	// Migrate the models to the database.
	database.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(&model.Board{}, &model.Task{})
		if err != nil {
			panic(err.Error())
		}
		return nil
	})
}

func main() {
	app := fiber.New()
	setupRoutes(app)

	err := app.Listen(":3020")
	if err != nil {
		panic(err.Error())
	}
}
