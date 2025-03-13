package model

import (
	"gorm.io/gorm"
	"time"
)

type Activity struct {
	gorm.Model
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Location  string    `json:"location"`
	Category  string    `json:"category"`
	UserID    uint      `json:"user_id"`
	Upvote    int       `json:"upvote"`
	Img       string    `json:"img"`
	Deadline  time.Time `json:"deadline"`
	StartTime time.Time `json:"start_time"`
}
