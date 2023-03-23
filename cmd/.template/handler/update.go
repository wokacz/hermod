package handler

import "github.com/gofiber/fiber/v2"

// Update
// function is used to modify existing records that exist in the database.
func Update(ctx *fiber.Ctx) (err error) {
	// id := ctx.Params("id")

	return ctx.Status(fiber.StatusOK).JSON(nil)
}
