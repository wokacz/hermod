package model

type Role struct {
	Model
	Name       string      `json:"name,omitempty" gorm:"uniqueIndex"`
	Privileges []Privilege `json:"privileges,omitempty" gorm:"many2many:role_privileges"`
}

/* Methods */

func (receiver Role) IsSuperUser() bool {
	return receiver.Name == "SuperUser"
}
