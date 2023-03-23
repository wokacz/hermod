package model

// Role is a model for a role. A role is a collection of privileges.
type Role struct {
	Model
	// Name is the name of the role. It is unique.
	Name string `json:"name,omitempty" gorm:"uniqueIndex"`
	// Privileges is a list of privileges that the role has.
	Privileges []Privilege `json:"privileges,omitempty" gorm:"many2many:role_privileges"`
}

// IsSuperUser returns true if the role is a super user.
func (r Role) IsSuperUser() bool {
	return r.Name == "super_user"
}
