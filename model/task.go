package model

import "github.com/google/uuid"

type Task struct {
	Model
	// Title is the title of the task.
	Title string `json:"title" gorm:"size:2048;not null" validate:"required,max=2048"`
	// Description is the description of the task.
	Description string `json:"description" gorm:"not null"`
	// CreatedBy is the ID of the user that created the task.
	CreatedBy uuid.UUID `json:"createdBy" gorm:"not null"`
	// AssignedTo is the ID of the user that the task is assigned to.
	AssignedTo uuid.UUID `json:"assignedTo"`
	// BoardID is the ID of the board that the task is in.
	BoardID uuid.UUID `json:"boardId" gorm:"not null" validate:"required"`
}
