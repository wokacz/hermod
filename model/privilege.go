package model

// Privilege is a model for a privilege. A privilege is a permission to do something.
type Privilege struct {
	Model
	// Name is the name of the privilege. It is unique.
	Name string `json:"name,omitempty" gorm:"uniqueIndex"`
	// Description is a description of the privilege.
	Description string `json:"description,omitempty"`
}
