package model

// Attachment is a model for an attachment. An attachment is a file that is
// attached to a task.
type Attachment struct {
	Model
	// Path is the path to the file.
	Path string `json:"path" gorm:"size:128;not null" validate:"required"`
	// Name is the name of the file.
	Name string `json:"name" gorm:"size:35;not null" validate:"required"`
	// Type is the MIME type of the file.
	Type string `json:"type" gorm:"size:32;not null" validate:"required"`
	// Size is the size of the file in bytes.
	Size int `json:"size" gorm:"not null"`
}
