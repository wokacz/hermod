package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wokacz/hermod/model"
	"github.com/wokacz/hermod/pkg/argon2"
	"github.com/wokacz/hermod/pkg/database"
	"github.com/wokacz/hermod/pkg/jwt"
	"github.com/wokacz/hermod/pkg/validation"
)

// SignIn signs in a user. It returns a JWT token.
func SignIn(ctx *fiber.Ctx) (err error) {
	// Parse the login credentials.
	var loginCredentials model.LoginCredentials
	_ = ctx.BodyParser(&loginCredentials)
	// Validate the login credentials. If the validation fails, return a bad request error.
	errors := validation.ValidateStruct(loginCredentials)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	// Get the user from the database.
	var user model.User
	data := database.DB.Where("user_name = ?", loginCredentials.UserName).First(&user)
	// If the user does not exist, return a not found error.
	if data.RowsAffected == 0 {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	// Compare the password. If the comparison fails, return an internal server error.
	match, err := argon2.Compare(loginCredentials.Password, user.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	// If the password does not match, return an unauthorized error.
	if !match {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	// Generate a JWT token for the user.
	token, err := jwt.Generate(&user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	// Return the token.
	return ctx.Status(fiber.StatusOK).SendString(token)
}
