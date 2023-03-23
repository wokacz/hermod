package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wokacz/hermod/model"
	"github.com/wokacz/hermod/pkg/argon2"
	"github.com/wokacz/hermod/pkg/database"
	"github.com/wokacz/hermod/pkg/jwt"
	"github.com/wokacz/hermod/pkg/validation"
	"gorm.io/gorm"
)

// SignUp signs up a user. It returns a JWT token.
func SignUp(ctx *fiber.Ctx) error {
	user := new(model.User)
	_ = ctx.BodyParser(user)
	// Validate the user. If the validation fails, return a bad request error.
	errors := validation.ValidateStruct(user)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	// Hash the password. If the hashing fails, return an internal server error.
	hash, err := argon2.Hash(user.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	// Set the password to the hashed password.
	user.Password = hash
	// Begin a transaction.
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		var role *model.Role
		// Get the role "user" or create it if it does not exist.
		tx.Model(role).FirstOrCreate(&role, model.Role{Name: "user"})
		// Add the role to the user.
		user.Roles = append(user.Roles, *role)
		// Create the user. If the user already exists, the transaction will be rolled back.
		err = tx.Create(user).Error
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	// Generate a JWT token for the user.
	token, err := jwt.Generate(user)
	if err != nil {
		return err
	}
	// Return the token.
	return ctx.Status(fiber.StatusCreated).SendString(token)
}
