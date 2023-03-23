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
	var user model.User
	_ = ctx.BodyParser(user)

	errors := validation.ValidateStruct(user)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	err = database.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(user).Error
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(nil)
}
