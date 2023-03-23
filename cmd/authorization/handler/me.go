package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wokacz/hermod/model"
	"github.com/wokacz/hermod/pkg/database"
)

// Me returns the user's data.
// It requires a valid JWT token.
func Me(ctx *fiber.Ctx) error {
	// Get the subject from the JWT token.
	subject := ctx.Locals("subject")
	var user *model.User
	// Get the user from the database.
	data := database.DB.Where("id = ?", subject).Preload("Roles").First(&user)
	// If the user does not exist, return a not found error.
	if data.RowsAffected == 0 {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	// Return the user.
	return ctx.Status(fiber.StatusOK).JSON(*user)
}
