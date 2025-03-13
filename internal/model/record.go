package model

import "gorm.io/gorm"

type Record struct {
	gorm.Model
	UserID     uint `json:"user_id"`
	ActivityID uint `json:"activity_id"`
}
