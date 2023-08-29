package model

import (
	resp "WheelChair-tiktok/model/response"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	AuthorID      uint   `gorm:"column:author_id"`
	PlayUrl       string `gorm:"column:play_url;unique"`
	CoverUrl      string `gorm:"column:cover_url;unique"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	Title         string `gorm:"column:title"`
}

func (*Video) TableName() string {
	return "video"
}

func (v *Video) ToResponse(isFavorite bool, author resp.User) resp.Video {

	return resp.Video{
		ID:            int64(v.ID),
		Author:        author,
		PlayURL:       v.PlayUrl,
		CoverURL:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    isFavorite,
		Title:         v.Title,
	}
}
