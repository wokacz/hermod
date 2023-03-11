package main

import (
	"github.com/wokacz/hermod/model"
	"github.com/wokacz/hermod/pkg/database"
	"gorm.io/gorm"
)

func main() {
	err := database.Init()
	if err != nil {
		panic(err.Error())
	}
	err = database.DB.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.AutoMigrate(
			&model.Privilege{},
			&model.Role{},
			&model.User{},
		)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		panic(err.Error())
	}
}
