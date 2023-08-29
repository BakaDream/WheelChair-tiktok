package response

type Feed struct {
	StatusCode int32   `json:"status_code"` //状态码 0 成功 其他失败
	StatusMsg  string  `json:"status_msg"`  // 状态信息
	VideoList  []Video `json:"video_list"`  //视频列表
	NextTime   int64   `json:"next_time"`   //本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

type Video struct {
	ID            int64  `json:"id"`             // 视频唯一标识
	Author        User   `json:"author"`         // 视频作者信息
	PlayURL       string `json:"play_url"`       //播放地址
	CoverURL      string `json:"cover_url"`      //封面地址
	FavoriteCount int64  `json:"favorite_count"` //点赞总数
	CommentCount  int64  `json:"comment_count"`  //评论总数
	IsFavorite    bool   `json:"is_favorite"`    //true-已点赞，false-未点赞
	Title         string `json:"title"`          //视频标题
}
type PublishList struct {
	StatusCode int32   `json:"status_code"` //状态码
	StatusMsg  string  `json:"status_msg"`  //状态信息
	VideoList  []Video `json:"video_list"`  // 视频列表
}
