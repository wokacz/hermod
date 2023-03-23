package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wokacz/hermod/pkg/database"
)

func setupRoutes(app *fiber.App) {
	app.Get("/__health", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
}

func init() {
	err := database.Init()
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	app := fiber.New()
	setupRoutes(app)

	err := app.Listen("")
	if err != nil {
		panic(err.Error())
	}
}
