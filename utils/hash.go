package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"path/filepath"
)

func HashFileName(fileName string) string {
	// 获取原始文件的拓展名
	extension := filepath.Ext(fileName)

	// 创建SHA256哈希对象
	hash := sha256.New()

	// 将文件名（不包括拓展名）转换为字节数组并进行哈希计算
	hash.Write([]byte(fileName[:len(fileName)-len(extension)]))

	// 获取哈希值并转换为十六进制字符串
	hashedBytes := hash.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)

	// 构建新的文件名，添加原始拓展名
	newFileName := hashedString + extension

	return newFileName
}
