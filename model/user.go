package model

import (
	resp "WheelChair-tiktok/model/response"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName        string `gorm:"unique;column:user_name"`                                            // 用户名称
	IP              string `gorm:"column:ip"`                                                          // 用户IP
	Password        string `gorm:"column:password"`                                                    // 密码
	FollowCount     int64  `gorm:"column:follow_count"`                                                // 关注总数
	FollowerCount   int64  `gorm:"column:follower_count"`                                              // 粉丝总数
	Signature       string `gorm:"default:'这个人很懒，没有简历哦~';column:signature"`                            // 个人简介
	Avatar          string `gorm:"column:avatar;default:'https://s1.ax1x.com/2023/08/30/pPdv8BR.png'"` // 用户头像
	BackgroundImage string `gorm:"column:background_image"`                                            // 用户个人页顶部大图
	TotalFavorited  int64  `gorm:"default:0;column:total_favorited"`                                   // 获赞数量，默认值为0
	WorkCount       int64  `gorm:"default:0;column:work_count"`                                        // 作品数量，默认值为0
	FavoriteCount   int64  `gorm:"default:0;column:favorite_count"`                                    // 点赞数量，默认值为0
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
