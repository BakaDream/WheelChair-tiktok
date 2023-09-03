package service

import (
	m "WheelChair-tiktok/model"
	"time"
)

func GetLastVideoTime() (time.Time, error) {
	var video m.Video
	err := m.DB.Last(&video).Error
	if err != nil {
		return time.Time{}, err
	}
	return video.CreatedAt, nil
}

func GetVideoList(lateTime time.Time) ([]m.Video, error) {
	var videos []m.Video
	err := m.DB.Where("created_at <= ?", lateTime).Limit(15).Find(&videos).Error
	if err != nil {
		return nil, err
	}

	return videos, nil
}
