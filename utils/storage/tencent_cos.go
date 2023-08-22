package storage

import (
	"WheelChair-tiktok/global"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
	"mime/multipart"
	"net/http"
	"net/url"
)

type TencentCos struct {
	CosUrl    string
	SecretId  string
	SecretKey string
}

// UploadFile 腾讯云cos上传，传入文件对象 返回云端路径
func (c *TencentCos) UploadFile(file *multipart.FileHeader) (string, error) {
	client := c.newClient()
	f, err := file.Open()
	if err != nil {
		global.Logger.Error("Fail", zap.Error(err))
		return "", err
	}
	defer f.Close()
	_, err = client.Object.Put(context.Background(), file.Filename, f, nil)
	ourl := client.Object.GetObjectURL(file.Filename)
	return ourl.String(), nil
}

func (c *TencentCos) newClient() *cos.Client {
	u, _ := url.Parse(c.CosUrl)
	// 用于 Get Service 查询，默认全地域 service.cos.myqcloud.com
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
