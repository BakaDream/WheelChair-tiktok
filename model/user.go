package model

import (
	resp "WheelChair-tiktok/model/response"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName        string `gorm:"unique" json:"user_name"` // 用户名称
	IP              string `json:"-"`                       // 用户IP
	Password        string `json:"-"`
	FollowCount     int64  `json:"follow_count"`                            // 关注总数
	FollowerCount   int64  `json:"follower_count"`                          // 粉丝总数
	Signature       string `gorm:"default:'这个人很懒，没有简历哦~'" json:"signature"` //个人简介
	Avatar          string `json:"avatar"`                                  //用户头像
	BackgroundImage string `json:"background_image"`                        //用户个人页顶部大图
	TotalFavorited  int64  `json:"total_favorited"`                         //获赞数量
	WorkCount       int64  `json:"work_count"`                              //作品数量
	FavoriteCount   int64  `json:"favorite_count"`                          //点赞数量
}

func (*User) TableName() string {
	return "user"
}

func (u *User) ToResponse(isFollow bool) resp.User {
	return resp.User{
		ID:              int64(u.ID),
		Name:            u.UserName,
		FollowCount:     u.FollowCount,
		FollowerCount:   u.FollowerCount,
		IsFollow:        isFollow,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		TotalFavorited:  u.TotalFavorited,
		WorkCount:       u.WorkCount,
		FavoriteCount:   u.FavoriteCount,
	}
}
