package storage

import (
	"WheelChair-tiktok/utils"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type Local struct {
	Static string
}

func (l *Local) UploadFile(file io.Reader, filePath string) (string, error) {
	dst := "public/" + filePath
	if err := os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return "", err
	}

	// 创建目标文件以供写入
	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close() // 在函数结束时关闭文件

	// 将上传文件的内容复制到目标文件
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}
	return l.Static + "/" + dst, nil
}

func (l *Local) GetSnapshot(videoFile *multipart.FileHeader) (string, error) {
	videoPath := "./public/" + "videos/" + videoFile.Filename
	name := utils.GetBaseNameWithoutExtension(videoFile.Filename)
	picName := name + ".jpg"
	picPath := "./public/" + "cover/" + picName
	if err := os.MkdirAll(filepath.Dir(picPath), 0750); err != nil {
		return "", err
	}
	err := utils.GetFirstFrame(videoPath, picPath)
	if err != nil {
		return "", err
	}
	picUrl := l.Static + "/public/cover/" + picName
	return picUrl, nil
}
