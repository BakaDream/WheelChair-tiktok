package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID  uint   `gorm:"column:user_id"`  //评论者id
	VideoID uint   `gorm:"column:video_id"` //视频id
	Content string `gorm:"column:content"`  //内容
}
