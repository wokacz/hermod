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

func main() {
	var err error

	err = database.Init()
	if err != nil {
		panic(err.Error())
	}

	app := fiber.New()
	setupRoutes(app)

	err = app.Listen(":5000")
	if err != nil {
		panic(err.Error())
	}
}
