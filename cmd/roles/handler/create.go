package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wokacz/hermod/model"
	"github.com/wokacz/hermod/pkg/database"
	"github.com/wokacz/hermod/pkg/validation"
	"gorm.io/gorm"
)

// Create
// function is used to create a new record in the database.
func Create(ctx *fiber.Ctx) (err error) {
	var role model.Role
	_ = ctx.BodyParser(role)
	errors := validation.ValidateStruct(role)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	err = database.DB.Transaction(func(tx *gorm.DB) error {
		tx.Create(role)
		return nil
	})
	if err != nil {
		return err
	}
	// Return the created role.
	return ctx.Status(fiber.StatusCreated).JSON(role)
}
