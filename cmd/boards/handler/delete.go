package handler

import "github.com/gofiber/fiber/v2"

// Delete
// function is used to remove records from a database that is no longer needed.
func Delete(ctx *fiber.Ctx) (err error) {
	// id := ctx.Params("id")

	return ctx.SendStatus(fiber.StatusNoContent)
}
