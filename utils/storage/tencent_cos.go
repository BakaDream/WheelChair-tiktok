package Storage

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

type TencentCos struct{}

// UploadFile 腾讯云cos上传，传入文件对象 返回云端路径
func (*TencentCos) UploadFile(file *multipart.FileHeader) (string, error) {
	client := newClient()
	f, err := file.Open()
	if err != nil {
		log.Fatal()
	}
	defer f.Close()
	_, err = client.Object.Put(context.Background(), file.Filename, f, nil)
	ourl := client.Object.GetObjectURL(file.Filename)
	return ourl.String(), nil
}

func newClient() *cos.Client {
	u, _ := url.Parse(os.Getenv(os.Getenv("Tencent_COS_URL")))
	// 用于 Get Service 查询，默认全地域 service.cos.myqcloud.com
	baseURL := &cos.BaseURL{BucketURL: u}
	// 1.永久密钥
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("Tencent_COS_SecretID"),  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: os.Getenv("Tencent_COS_SecretKey"), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	return client
}
