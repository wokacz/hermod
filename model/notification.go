package model

import "github.com/google/uuid"

// Notification is a model for a notification. A notification is a message that is sent to a user.
type Notification struct {
	Model
	// UserID is the ID of the user that the notification is sent to.
	UserID uuid.UUID `json:"userId" gorm:"not null" validate:"required"`
	// Message is the message of the notification.
	Message string `json:"message"`
	// Read is a boolean that indicates if the notification has been read.
	Read bool `json:"read"`
}
