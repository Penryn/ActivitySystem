package service

import (
	"activitySystem/internal/dao"

	"gorm.io/gorm"
)

var (
	d *dao.Dao
)

func Init(db *gorm.DB) {
	d = dao.New(db)
}
