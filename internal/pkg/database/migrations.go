package database

import (
	"activitySystem/internal/model"

	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
		model.Activity{},
		model.Record{},
	)

}
