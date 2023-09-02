package service

import (
	"WheelChair-tiktok/utils/storage"
	"mime/multipart"
)

/*
UploadVideo 上传视频
params: file c.FormFile 获取的对象
return videoPath, coverPath, error
*/
func UploadVideo(file *multipart.FileHeader) (string, string, error) {
	//打开文件
	f, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer f.Close()
	filePath := "videos/" + file.Filename
	playUrl, err := storage.Storage.UploadFile(f, filePath)
	if err != nil {
		return "", "", err
	}
	coverUrl, err := storage.Storage.GetSnapshot(file)
	if err != nil {
		return "", "", err
	}
	return playUrl, coverUrl, nil
}
