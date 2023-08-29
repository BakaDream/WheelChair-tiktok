package model

type Favorite struct {
	UserID  uint
	VideoID uint
}

func (Favorite) TableName() string {
	return "favorite"
}
