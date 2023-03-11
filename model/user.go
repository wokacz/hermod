package model

type User struct {
	Model
	FirstName string `json:"firstName,omitempty" validate:"required,max=32"`
	LastName  string `json:"lastName,omitempty" validate:"required,max=32"`
	Email     string `json:"email,omitempty" gorm:"uniqueIndex" validate:"required,email"`
	UserName  string `json:"userName,omitempty" gorm:"uniqueIndex" validate:"required,min=6,max=32"`
	Password  string `json:"password,omitempty" validate:"required,min=6,max=128"`
	Roles     []Role `json:"roles,omitempty" gorm:"many2many:user_roles"`
}
