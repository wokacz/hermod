package handler

import "github.com/gofiber/fiber/v2"

// TakeMany
// function is used to search and retrieve specific records in the table and read their values.
// Users may be able to find desired records using keywords, or by filtering the data based on customized criteria.
func TakeMany(ctx *fiber.Ctx) (err error) {
	return ctx.Status(fiber.StatusOK).JSON(nil)
}

// TakeOne
// It allows users to search and retrieve a specific record based on ID.
func TakeOne(ctx *fiber.Ctx) (err error) {
	// id := ctx.Params("id")

	return ctx.Status(fiber.StatusOK).JSON(nil)
}
