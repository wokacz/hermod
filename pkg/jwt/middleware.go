package jwt

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Config struct is used to define the configuration options for the middleware.
type Config struct {
	Secret string
}

// New function is used to create a new JWT middleware.
// The middleware will check for a valid JWT token in the Authorization header.
// If the token is valid, the subject claim will be added to the request context.
// If the token is invalid, a 401 status code will be returned.
func New(config Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Get the Authorization header.
		authHeader := ctx.Get("Authorization")
		// If the Authorization header is empty, return a 401 status code.
		if authHeader == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "authorization header is empty")
		}
		// Split the Authorization header into parts.
		authParts := strings.Split(authHeader, " ")
		// If the Authorization header does not contain two parts, return a 401 status code.
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			return ctx.Status(fiber.StatusUnauthorized).SendString("invalid authorization header")
		}
		// Parse the token.
		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != "HS256" {
				return nil, errors.New("invalid token algorithm")
			}
			return []byte(config.Secret), nil
		})
		// If the token is invalid, return a 401 status code.
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid token")
		}
		// If the token is invalid, return a 401 status code.
		if !token.Valid {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid token")
		}
		// Get the subject claim from the token.
		subject, err := token.Claims.GetSubject()
		// If the subject claim is empty, return a 401 status code.
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "something went wrong")
		}
		// Add the subject claim to the request context.
		ctx.Locals("subject", subject)
		// Call the next handler.
		return ctx.Next()
	}
}
