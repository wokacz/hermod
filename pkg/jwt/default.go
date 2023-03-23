package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/wokacz/hermod/model"
	"github.com/wokacz/hermod/pkg/env"
)

// Claims struct is used to define the claims that are added to the JWT token.
type Claims struct {
	jwt.RegisteredClaims
}

// Generate function is used to generate a JWT token.
// The token is signed using the JWT_SECRET environment variable.
func Generate(user *model.User) (token string, err error) {
	// Get the JWT secret from the environment.
	secret := env.Get("JWT_SECRET", "top-secret")
	// Create a new JWT token.
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// TODO: Add the user's roles to the token.
			Issuer:  env.Get("APP_NAME", "App"),
			Subject: user.ID.String(),
		},
	})
	// Sign the token using the JWT secret.
	token, err = jwtToken.SignedString([]byte(secret))
	return token, err
}
