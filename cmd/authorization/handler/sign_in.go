package handler

import "github.com/gofiber/fiber/v2"

func SignIn(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusUnauthorized)
}
