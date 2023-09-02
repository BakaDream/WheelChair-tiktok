package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	UserID  uint `gorm:"column:user_id;index"`
	VideoID uint `gorm:"column:video_id;index"`
}

func (Favorite) TableName() string {
	return "favorite"
}
