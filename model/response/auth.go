package response

type Auth struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}
