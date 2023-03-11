package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wokacz/hermod/internal/validator"
	"github.com/wokacz/hermod/model"
)

func SignUp(ctx *fiber.Ctx) error {
	user := new(model.User)
	_ = ctx.BodyParser(user)
	errors := validator.ValidateStruct(user)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return ctx.SendStatus(fiber.StatusCreated)
}
