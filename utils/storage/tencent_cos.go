package storage

import (
	"WheelChair-tiktok/utils"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

type TencentCos struct {
	CosUrl    string
	SecretId  string
	SecretKey string
}

// UploadFile 传入io.Reader ,filepath 返回云端路径 ,err 记得在外部要关闭io.Reader
func (c *TencentCos) UploadFile(file io.Reader, filePath string) (string, error) {
	client := c.newClient()
	_, err := client.Object.Put(context.Background(), filePath, file, nil)
	if err != nil {
		return "", err
	}
	fileUrl := c.CosUrl + "/" + filePath
	return fileUrl, nil
}
func (c *TencentCos) newClient() *cos.Client {
	u, _ := url.Parse(c.CosUrl)
	baseURL := &cos.BaseURL{BucketURL: u}
	// 1.永久密钥
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  c.SecretId,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: c.SecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	return client
}

func (c *TencentCos) GetSnapshot(videoFile *multipart.FileHeader) (string, error) {
	name := utils.GetBaseNameWithoutExtension(videoFile.Filename)
	picName := name + ".jpg"
	picPath := "cover/" + picName
	f, err := videoFile.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()
	// 创建临时视频文件
	tempVideo, err := os.CreateTemp("", videoFile.Filename)
	if err != nil {
		// Handle ERROR
		return "", err
	}
	defer os.Remove(tempVideo.Name()) // 清理临时文件
	defer tempVideo.Close()           // 关闭临时文件
	// 将响应体写入临时文件
	_, err = io.Copy(tempVideo, f)
	if err != nil {
		// Handle ERROR
		return "", err
	}

	tempPic, err := os.CreateTemp("", picName)
	if err != nil {
		// Handle ERROR
		return "", err
	}
	defer os.Remove(tempPic.Name()) // 清理临时文件
	defer tempPic.Close()
	//生成封面
	utils.GetFirstFrame(tempVideo.Name(), tempPic.Name())

	//上传临时文件
	picUrl, err := c.UploadFile(tempPic, picPath)
	if err != nil {
		return "", err
	}
	return picUrl, nil
}
