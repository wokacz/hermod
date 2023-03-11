package model

type Privilege struct {
	Model
	Name        string `json:"name,omitempty" gorm:"uniqueIndex"`
	Description string `json:"description,omitempty"`
}
