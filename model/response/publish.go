package response

type Publish struct {
	StatusCode int32  `json:"status_code,omitempty"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
