package acl

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wokacz/hermod/model"
	"github.com/wokacz/hermod/pkg/database"
)

type Config struct {
	Roles []string
}

// New function is used to create a new ACL middleware.
// The middleware will check if the user has the required roles.
func New(config Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		subject := ctx.Locals("subject")

		var user model.User

		data := database.DB.Where("id = ?", subject).Preload("Roles").Select("id").First(&user)
		if data.RowsAffected == 0 {
			return ctx.SendStatus(fiber.StatusNotFound)
		}

		hasRole := false

		for _, userRole := range user.Roles {
			for _, allowedRole := range config.Roles {
				if userRole.IsSuperUser() || userRole.Name == allowedRole {
					hasRole = true
					break
				}
			}
		}

		if hasRole {
			return ctx.Next()
		} else {
			return ctx.SendStatus(fiber.StatusForbidden)
		}
	}
}
