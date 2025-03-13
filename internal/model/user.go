package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	StuID    string `json:"stu_id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Avatar   string `json:"avatar"`
	Profile  string `json:"profile"`
}
