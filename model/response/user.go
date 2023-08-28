package response

import m "WheelChair-tiktok/model"

type Register struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

type Login struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

type UserInfo struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	User       m.User `json:"user"`
}
