package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/wokacz/hermod/cmd/authorization/handler"
	"github.com/wokacz/hermod/pkg/database"
)

func setupRoutes(app *fiber.App) {
	app.Get("/__health", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
	app.Post("/sign-in", handler.SignIn)
	app.Post("/sign-up", handler.SignUp)
}

func main() {
	var err error

	err = database.Init()
	if err != nil {
		panic(err.Error())
	}

	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	err = app.Listen(":3000")
	if err != nil {
		panic(err.Error())
	}
}
