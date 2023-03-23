package handler

import "github.com/gofiber/fiber/v2"

// Create
// function is used to create a record in the database.
func Create(ctx *fiber.Ctx) (err error) {

	return ctx.Status(fiber.StatusCreated).JSON(nil)
}
