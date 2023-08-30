package service

import (
	m "WheelChair-tiktok/model"
	"gorm.io/gorm"
)

func AddPublish(playUrl string, coverUrl string, title string, uid uint) error {
	// 在数据库里新建一条视频记录
	video := m.Video{
		AuthorID: uid,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    title,
	}
	err := m.DB.Create(&video).Error
	if err != nil {
		return err
	}
	//对应用户workCount 加1
	var user m.User
	err = m.DB.Model(&user).Where("ID = ?", uid).Update("work_count", gorm.Expr("work_count + ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPublishList(authorID uint) ([]m.Video, error) {
	var videos []m.Video
	// Find all videos with the specified AuthorID
	err := m.DB.Where("author_id = ?", authorID).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	if len(videos) == 0 {
		return nil, nil
	}

	return videos, nil
}
