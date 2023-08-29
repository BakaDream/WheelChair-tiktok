package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID  uint   `gorm:"user_id" json:"-"`
	VideoID uint   `json:"video_id"`
	Content string `gorm:"not null" json:"content"`
}
