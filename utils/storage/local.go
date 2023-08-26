package storage

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type Local struct {
	Static string
}

func (l *Local) UploadFile(file *multipart.FileHeader) (string, error) {
	dst := "public/" + "videos/" + file.Filename
	// 打开上传的文件
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close() // 在函数结束时关闭文件，以防止资源泄漏

	// 创建目标目录，如果不存在的话
	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return "", err
	}

	// 创建目标文件以供写入
	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close() // 在函数结束时关闭文件

	// 将上传文件的内容复制到目标文件
	_, err = io.Copy(out, f)
	if err != nil {
		return "", err
	}
	return l.Static + "/" + dst, nil
}
