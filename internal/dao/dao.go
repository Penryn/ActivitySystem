package dao

import (
	"gorm.io/gorm"
)

type Dao struct {
	orm *gorm.DB
}

func New(orm *gorm.DB) *Dao {
	return &Dao{orm: orm}
}
