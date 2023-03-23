package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/wokacz/hermod/cmd/authorization/handler"
	"github.com/wokacz/hermod/pkg/argon2"
	"github.com/wokacz/hermod/pkg/database"
	"github.com/wokacz/hermod/pkg/env"
	"github.com/wokacz/hermod/pkg/jwt"
)

// setupRoutes sets up the routes.
func setupRoutes(app *fiber.App) {
	// Health check.
	app.Get("/__health", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	app.Post("/sign-in", handler.SignIn)
	app.Post("/sign-up", handler.SignUp)

	// JWT middleware.
	secret := env.Get("JWT_SECRET", "top-secret")
	app.Use(jwt.New(jwt.Config{
		Secret: secret,
	}))

	app.Get("/me", handler.Me)
}

func init() {
	err := database.Init()
	if err != nil {
		panic(err.Error())
	}
	argon2.Init()
}

func main() {
	var err error
	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)
	err = app.Listen(":3000")
	if err != nil {
		panic(err.Error())
	}
}
