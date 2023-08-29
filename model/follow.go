package model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	FollowedID uint //被关注人的ID
	FollowerID uint //关注者ID
}

func (Follow) TableName() string {
	return "follow"
}
