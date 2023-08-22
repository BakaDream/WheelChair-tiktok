package storage

import (
	"WheelChair-tiktok/global"
	"mime/multipart"
	"os"
)

type Store interface {
	UploadFile(file *multipart.FileHeader) (string, error)
}

func Init() {
	switch os.Getenv("STORAGE_TYPE") {
	case "TencentCOS":
		global.Storage = &TencentCos{
			CosUrl:    os.Getenv("Tencent_COS_URL"),
			SecretId:  os.Getenv("Tencent_COS_SecretKey"),
			SecretKey: os.Getenv("Tencent_COS_SecretKey"),
		}
	case "Local":
		global.Storage = &Local{}
	default:
		global.Logger.Fatal("Can't Load Storage Config, Please Check env")

	}

}
