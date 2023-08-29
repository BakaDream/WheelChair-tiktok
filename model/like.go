package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	UserID  uint
	VideoID uint
}

func (Favorite) TableName() string {
	return "favorite"
}
