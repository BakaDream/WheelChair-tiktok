package response

// FavoriteAction 赞操作 返回
type FavoriteAction struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

// FavoriteList 喜欢列表
type FavoriteList struct {
	StatusCode int32   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 用户点赞视频列表
}
