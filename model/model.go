package model

import "time"

type Response struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
	User       User   `json:"user"`                 // 用户信息
}

type Video struct {
	Id            int64     `json:"id,omitempty"`
	AuthorId      int64     `json:"authorid"`
	PlayUrl       string    `json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count,omitempty"`
	CommentCount  int64     `json:"comment_count,omitempty"`
	IsFavorite    bool      `json:"is_favorite,omitempty"`
	Title         string    `json:"title,omitempty"`
	PublishTime   time.Time `json:"publish_time"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id              int64  `json:"id,omitempty"`             // 用户id
	Name            string `json:"name,omitempty"`           // 用户名称
	FollowCount     int64  `json:"follow_count,omitempty"`   // 关注总数
	FollowerCount   int64  `json:"follower_count,omitempty"` // 粉丝总数
	IsFollow        bool   `json:"is_follow,omitempty"`      // true-已关注，false-未关注
	Signature       string `json:"signature"`                //个人简介
	Avatar          string `json:"avatar"`                   //用户头像
	BackgroundImage string `json:"background_image"`         //用户个人页顶部大图
	TotalFavorited  string `json:"total_favorited"`          //获赞数量
	WorkCount       int64  `json:"work_count"`               //作品数量
	FavoriteCount   int64  `json:"favorite_count"`           //点赞数量
}

type Message struct {
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
	FromUserID int64  `json:"from_user_id"`
	ID         int64  `json:"id"`
	ToUserID   int64  `json:"to_user_id"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
