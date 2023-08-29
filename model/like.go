package model

type Like struct {
	UserID  uint
	VideoID uint
}

func (Like) TableName() string {
	return "like"
}
