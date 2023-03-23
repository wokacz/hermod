package model

// LoginCredentials struct is used to validate login credentials.
type LoginCredentials struct {
	// UserName is the username of the user.
	UserName string `json:"userName" validate:"required,min=6,max=32"`
	// Password is the password of the user.
	Password string `json:"password" validate:"required,min=6,max=128"`
}
