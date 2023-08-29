package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	UserID        uint   `gorm:"user_id" json:"-"`
	PlayUrl       string `gorm:"unique" json:"play_url"`
	CoverUrl      string `gorm:"unique" json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	Title         string `gorm:"not null" json:"title"`
}

func (Video) TableName() string {
	return "video"
}
