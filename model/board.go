package model

type Board struct {
	Model
	// Name is the name of the board.
	Name string `json:"name" gorm:"not null"`
	// Tasks []Task `json:"tasks" gorm:"foreignKey:BoardTasks;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
