package model

import "strings"

// User struct is used to represent a user.
type User struct {
	Model
	// UserName is the username of the user. It is unique and cannot be null.
	UserName string `json:"userName" gorm:"size:32;uniqueIndex;not null" validate:"required,min=6,max=32"`
	// Password is the password of the user. It cannot be null.
	Password string `json:"-" gorm:"size:128;not null" validate:"required,min=6,max=128"`
	// FirstName is the first name of the user. It cannot be null.
	FirstName string `json:"firstName" validate:"required,max=32"`
	// LastName is the last name of the user. It cannot be null.
	LastName string `json:"lastName" validate:"required,max=32"`
	// Email is the email of the user. It is unique and cannot be null.
	Email string `json:"email" gorm:"uniqueIndex" validate:"required,email"`
	// Roles is a many-to-many relationship with the Role model.
	Roles []Role `json:"roles,omitempty" gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Notifications is a many-to-many relationship with the Notification model.
	Notifications []Notification `json:"notifications,omitempty" gorm:"many2many:user_notifications;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// BeforeSave is a hook that is executed before the user is saved.
func (u User) BeforeSave() {
	// Convert the email and username to lowercase.
	u.Email = strings.ToLower(u.Email)
	u.UserName = strings.ToLower(u.UserName)
}
