package service

import (
	m "WheelChair-tiktok/model"
	"gorm.io/gorm"
)

func Favorite(videoID uint, userID uint) error {
	// todo:原子操作
	favorite := m.Favorite{
		UserID:  userID,
		VideoID: videoID,
	}

	// 检查是否已存在软删除记录
	var existingFavorite m.Favorite
	err := m.DB.Unscoped().Where("user_id = ? AND video_id = ?", userID, videoID).First(&existingFavorite).Error
	if err == nil {
		// 取消软删除
		err = m.DB.Unscoped().Model(&existingFavorite).Update("deleted_at", gorm.Expr("NULL")).Error
		if err != nil {
			return err
		}
	} else {
		// 创建点赞记录
		err = m.DB.Create(&favorite).Error
		if err != nil {
			return err
		}
	}

	var video m.Video
	//先更新video结构体的信息
	err = m.DB.Model(&video).Where("ID = ?", videoID).First(&video).Error
	if err != nil {
		return err
	}
	//更新视频的FavoriteCount 获赞计数
	err = m.DB.Model(&video).Where("ID = ?", videoID).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	if err != nil {
		return err
	}
	//更新用户的FavoriteCount 的点赞计数
	var user m.User
	err = m.DB.Model(&user).Where("ID = ?", userID).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	if err != nil {
		return err
	}
	//更新视频作者的TotalFavorited 的获赞计数
	var author m.User
	err = m.DB.Model(&author).Where("ID = ?", video.AuthorID).Update("total_favorited", gorm.Expr("total_favorited + ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func UnFavorite(videoID uint, userID uint) error {
	// todo:原子操作
	var favorite m.Favorite
	err := m.DB.Where("user_id = ? AND video_id = ?", userID, videoID).First(&favorite).Error
	if err != nil {
		return err
	}
	// 执行软删除
	err = m.DB.Delete(&favorite).Error
	if err != nil {
		return err
	}
	var video m.Video
	//先更新video结构体的信息
	err = m.DB.Model(&video).Where("ID = ?", videoID).First(&video).Error
	if err != nil {
		return err
	}
	//更新视频的FavoriteCount 获赞计数
	err = m.DB.Model(&video).Where("ID = ?", videoID).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
	if err != nil {
		return err
	}
	//更新用户的FavoriteCount 的点赞计数
	var user m.User
	err = m.DB.Model(&user).Where("ID = ?", userID).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
	if err != nil {
		return err
	}
	//更新视频作者的TotalFavorited 的获赞记录
	var author m.User
	err = m.DB.Model(&author).Where("ID = ?", video.AuthorID).Update("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func IsFavorite(userID uint, videoID uint) bool {
	var favorite m.Favorite
	//[2.186ms] [rows:0] SELECT * FROM `favorite` WHERE (user_id = 4 AND video_id = 9) AND `favorite`.`deleted_at` IS NULL ORDER BY `favorite`.`id` LIMIT 1
	err := m.DB.Where("user_id = ? AND video_id = ?", userID, videoID).First(&favorite).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		} else {
			return false
		}
	}
	return true
}
