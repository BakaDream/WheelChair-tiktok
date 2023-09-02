package model

import (
	resp "WheelChair-tiktok/model/response"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint   `gorm:"column:user_id"`  //评论者id
	VideoID uint   `gorm:"column:video_id"` //视频id
	Content string `gorm:"column:content"`  //内容
}

func (c *Comment) ToResponse(user resp.User) resp.Comment {
	return resp.Comment{
		ID:         int64(c.ID),
		User:       &user,
		Content:    c.Content,
		CreateDate: c.CreatedAt.Format("01-02"),
	}
}
