package storage

import (
	"WheelChair-tiktok/logger"
	"io"
	"mime/multipart"
	"os"
)

var Storage Store

type Store interface {
	UploadFile(file io.Reader, filePath string) (string, error)
	GetSnapshot(videoFile *multipart.FileHeader) (string, error)
}

func Init() Store {
	switch os.Getenv("STORAGE_TYPE") {
	case "TencentCOS":
		Storage = &TencentCos{
			CosUrl:    os.Getenv("TENCENT_COS_URL"),
			SecretId:  os.Getenv("TENCENT_COS_SECRET_ID"),
			SecretKey: os.Getenv("TENCENT_COS_SECRET_KEY"),
		}
	case "Local":
		Storage = &Local{
			Static: os.Getenv("LOCAL_STATIC_URL"),
		}
	default:
		logger.Logger.Fatal("Storage type has some err")
	}
	return nil
}
