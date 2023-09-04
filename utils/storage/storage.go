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

func Init() {
	switch os.Getenv("STORAGE_TYPE") {
	case "TencentCOS":
		Storage = &TencentCos{
			CosUrl:    os.Getenv("TENCENT_COS_URL"),
			SecretId:  os.Getenv("TENCENT_COS_SECRET_ID"),
			SecretKey: os.Getenv("TENCENT_COS_SECRET_KEY"),
		}
	case "Local":
		Storage = &Local{
			Static: os.Getenv("STATIC_URL"),
		}
	//case "MinIO":
	//	var useSSL = false
	//	if os.Getenv("MINIO_USESSL") == "true" {
	//		useSSL = true
	//	}
	//	Storage = &MinIO{
	//		Endpoint:        os.Getenv("MINIO_ENDPOINT"),
	//		AccessKeyID:     os.Getenv("MINIO_ACCESSKEYID"),
	//		SecretAccessKey: os.Getenv("MINIO_SECRECTACCESSKEY"),
	//		UseSSL:          useSSL,
	//		BucketName:      os.Getenv("MINIO_BUCKETNAME"),
	//	}

	default:
		logger.Logger.Fatal("Storage type has some err")
	}
	return
}
