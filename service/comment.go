package service

import (
	m "WheelChair-tiktok/model"
	"gorm.io/gorm"
)

func PublishComment(videoID uint, userID uint, commentText string) (*m.Comment, error) {
	// 新建数据库事务
	tx := m.DB.Begin()
	newComment := m.Comment{
		UserID:  userID,
		VideoID: videoID,
		Content: commentText,
	}
	//新增一条记录
	err := tx.Create(&newComment).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//更新video的CommentCount
	var video m.Video
	err = tx.Model(&video).Where("id = ?", videoID).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//提交事务
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &newComment, nil
}

func DeleteComment(commentID uint) error {
	//开启事务
	tx := m.DB.Begin()
	var comment m.Comment
	err := tx.First(&comment, commentID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Delete(&comment).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//更新video的CommentCount
	var video m.Video
	err = tx.Model(&video).Where("id = ?", comment.VideoID).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func GetCommentList(videoID uint) ([]m.Comment, error) {
	var comments []m.Comment
	if err := m.DB.Where("video_id = ?", videoID).Find(&comments).Error; err != nil {
		return comments, err
	}
	//没错误，返回
	return comments, nil

}
