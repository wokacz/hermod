package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wokacz/hermod/model"
	"github.com/wokacz/hermod/pkg/database"
	"github.com/wokacz/hermod/pkg/validation"
	"gorm.io/gorm"
)

// Update function is used to modify existing records that exist in the database.
func Update(ctx *fiber.Ctx) (err error) {
	// Parse the user struct from the request body.
	var user model.User
	_ = ctx.BodyParser(user)
	// Validate the user struct fields and return any errors if they exist.
	errors := validation.ValidateStruct(user)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	// Get the ID from the request parameters.
	id := ctx.Params("id")
	// Get the user.
	var userRow model.User
	data := database.DB.Where("id = ?", id).First(userRow)
	// If the user does not exist, return a 404 status code.
	if data.RowsAffected == 0 {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	// Begin a transaction.
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		// Update the user record.
		err := tx.Model(&userRow).Updates(user).Error
		return err
	})
	if err != nil {
		return err
	}
	// Return a 200 status code.
	return ctx.Status(fiber.StatusOK).JSON(nil)
}
