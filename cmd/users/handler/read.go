package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wokacz/hermod/model"
	"github.com/wokacz/hermod/pkg/database"
)

// TakeMany function is used to retrieve all records from the database.
func TakeMany(ctx *fiber.Ctx) (err error) {
	//value := ctx.Locals("subject")
	var users []model.User
	database.DB.Omit("password").Find(&users)
	// Return the records.
	return ctx.Status(fiber.StatusOK).JSON(users)
}

// TakeOne function is used to retrieve a single record from the database.
func TakeOne(ctx *fiber.Ctx) (err error) {
	// Retrieve the ID from the URL.
	id := ctx.Params("id")
	// Create a new instance of the user model. This is used to store the record from the database.
	var user model.User
	// Retrieve the record from the database.
	data := database.DB.Omit("password").Where("id = ?", id).First(&user)
	// If the record does not exist, return a 404 status code.
	if data.RowsAffected == 0 {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	// Return the record.
	return ctx.Status(fiber.StatusOK).JSON(user)
}
