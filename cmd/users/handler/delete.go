package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wokacz/hermod/model"
	"github.com/wokacz/hermod/pkg/database"
)

// Delete
// function is used to remove records from a database that is no longer needed.
func Delete(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	var user model.User

	data := database.DB.Where("id = ?", id).First(user)
	if data.RowsAffected == 0 {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	
	return ctx.SendStatus(fiber.StatusNoContent)
}
