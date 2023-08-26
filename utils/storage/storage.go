package storage

import (
	"log"
	"mime/multipart"
	"os"
)

var Storage Store

type Store interface {
	UploadFile(file *multipart.FileHeader) (string, error)
}

func Init() Store {
	switch os.Getenv("STORAGE_TYPE") {
	case "TencentCOS":
		Storage = &TencentCos{
			CosUrl:    os.Getenv("TENCENT_COS_URL"),
			SecretId:  os.Getenv("Tencent_COS_SecretKey"),
			SecretKey: os.Getenv("Tencent_COS_SecretKey"),
		}
	case "Local":
		Storage = &Local{
			Static: os.Getenv("LOCAL_STATIC_URL"),
		}
	default:
		log.Fatal("Can't Load Storage Config, Please Check env")

	}
	return nil
}
