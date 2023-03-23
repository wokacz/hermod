package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Model is the base model for all models.
// It contains the ID, CreatedAt, UpdatedAt and DeletedAt fields.
// The ID field is a UUID and is automatically generated.
// The DeletedAt field is used for soft deletes.
type Model struct {
	// ID is the primary key for the model.
	ID uuid.UUID `json:"ID" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	// CreatedAt is the timestamp when the model was created.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt is the timestamp when the model was updated.
	UpdatedAt time.Time `json:"updatedAt"`
	// DeletedAt is the timestamp when the model was deleted.
	// If the model is not deleted, the value will be null.
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
}
