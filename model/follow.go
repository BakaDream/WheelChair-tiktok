package model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	FollowedID uint `gorm:"followed_id;index"` //被关注人的ID
	FollowerID uint `gorm:"follower_id;index"` //关注者ID
}

func (Follow) TableName() string {
	return "follow"
}
