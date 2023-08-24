package model

import (
	"gorm.io/gorm"
)

// ---------------数据库and信息 模型-----------------
// 使用自动迁移模型创建表
type Video struct {
	//Id            uint      `gorm:"primaryKey" json:"id"`
	gorm.Model
	UserID        uint   `json:"user_id"`
	PlayUrl       string `gorm:"unique" json:"play_url"`
	CoverUrl      string `gorm:"unique" json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	Title         string `gorm:"not null" json:"title"`
	IsFavorite    bool   `gorm:"-" json:"is_favorite"` //只在传递信息用到
}

type Comment struct {
	//Id            uint      `gorm:"primaryKey" json:"id"`
	gorm.Model
	UserID  uint   `json:"user_id"`
	Content string `gorm:"not null" json:"content"`
}

type User struct {
	//Id              uint      `gorm:"primaryKey" json:"id"`    // 用户id
	gorm.Model
	UserName        string `json:"user_name"` // 用户名称
	IP              string `json:"ip"`
	Password        string `json:"-"`
	FollowCount     int64  `json:"follow_count"`                            // 关注总数
	FollowerCount   int64  `json:"follower_count"`                          // 粉丝总数
	Signature       string `gorm:"default:'这个人很懒，没有简历哦~'" json:"signature"` //个人简介
	Avatar          string `json:"avatar"`                                  //用户头像
	BackgroundImage string `json:"background_image"`                        //用户个人页顶部大图
	TotalFavorited  string `json:"total_favorited"`                         //获赞数量
	WorkCount       int64  `json:"work_count"`                              //作品数量
	FavoriteCount   int64  `json:"favorite_count"`                          //点赞数量
	IsFollow        bool   `gorm:"-" json:"is_favorite"`                    //只在传递信息用到
	//CreateUser      time.Time `json:"create_user"`
}

// ---------------喜欢 模型------------------
type UserVideoLike struct {
	gorm.Model
	UserID  uint `gorm:"user_id"`
	VideoId uint `gorm:"video_id"`
}

// ---------------评论 模型------------------
type UserVideoComment struct {
	gorm.Model
	UserID    uint `gorm:"user_id"`
	VideoId   uint `gorm:"video_id"`
	CommentID uint `gorm:"comment_id"`
}

// ---------------关注 模型------------------
//
//	type UserFollow struct {
//		gorm.Model
//		UserId1 uint `json:"user_id_1"`
//		UserId2 uint `json:"user_id_2"`
//	}
//
// ---------------视频流 模型------------------
type FeedRequest struct {
	LatestTime int64  `json:"latestTime"`
	Token      string `json:"token"`
}

type FeedResponse struct {
	StatusCode int32   `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg,omitempty"` // 返回状态描述
	VideoList  []Video `json:"video_list"`
	NextTime   int64   `json:"next_time"` // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

/*
用户在移动应用程序中浏览视频内容时，向后端服务器发出请求，服务器响应这些请求并提供相应的视频数据。
前端应用程序会发送 FeedRequest 请求来获取视频列表，然后根据 FeedResponse 中的数据显示视频内容。
*/

// ---------------用户注册 模型------------------
type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserRegisterResponse struct {
	StatusCode int32  `json:"statusCode"`           // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
	UserID     uint   `json:"user_id"`
	Token      string `json:"token"`
}

// ---------------获取用户信息 模型------------------
type UserRequest struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserResponse struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
	UserID     uint   `json:"user"`
}

// ---------------视频发布 模型------------------
type PublishRequest struct {
	Data  []byte `json:"data"`
	Title string `json:"title"`
	Token string `json:"token"`
}

type PublishResponse struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
	Video      Video  `json:"video"`                //直接返回视频
}

// ---------------已发布视频 模型------------------
type UserPublishRequest struct {
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}
type UserPublishResponse struct {
	StatusCode  int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg   string `json:"status_msg,omitempty"` // 返回状态描述
	VideoListID []uint `json:"video_list_id"`
}

// ---------------赞 模型------------------
// 登录用户对视频的点赞和取消点赞操作。
type FavoriteActionRequest struct {
	VideoID    uint   `json:"video_id"`
	ActionType int32  `json:"action_type"` // 1-点赞，2-取消点赞
	Token      string `json:"token"`
}
type FavoriteActionResponse struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
}

// ---------------喜欢列表 模型------------------
type FavoriteListRequest struct {
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}
type FavoriteListResponse struct {
	StatusCode  int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg   string `json:"status_msg,omitempty"` // 返回状态描述
	VideoListID []uint `json:"video_list_id"`
}

// ---------------评论 模型------------------
type CommentActionRequest struct {
	VideoID     uint   `json:"video_id"`
	ActionType  int32  `json:"action_type"` //1-发布评论，2-删除评论
	CommentText string `json:"comment_text"`
	CommentId   int64  `json:"comment_id"`
	Token       string `json:"token"`
}
type CommentActionResponse struct {
	StatusCode int32   `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg,omitempty"` // 返回状态描述
	Comment    Comment `json:"comment"`              // 评论成功返回评论内容，不需要重新拉取整个列表
}

// ---------------评论列表 模型------------------
type CommentListRequest struct {
	VideoID uint   `json:"video_id"`
	Token   string `json:"token"`
}
type CommentListResponse struct {
	StatusCode    uint   `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg     string `json:"status_msg,omitempty"` // 返回状态描述
	CommentListID []uint `json:"comment_list_id"`
}

// ----------------错误信息 模型-------------------
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// ---------------消息 模型------------------------
//type Message struct {
//	Content    string `json:"content"`
//	CreateTime int64  `json:"create_time"`
//	FromUserID int64  `json:"from_user_id"`
//	ID         int64  `json:"id"`
//	ToUserID   int64  `json:"to_user_id"`
//}
//
//type MessageSendEvent struct {
//	UserID     int64  `json:"user_id,omitempty"`
//	ToUserID   int64  `json:"to_user_id,omitempty"`
//	MsgContent string `json:"msg_content,omitempty"`
//}
//
//type MessagePushEvent struct {
//	FromUserID int64  `json:"user_id,omitempty"`
//	MsgContent string `json:"msg_content,omitempty"`
//}
//
//type Response struct {
//	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
//	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
//	UserID     int64  `json:"user"`                 // 用户信息
//}
