package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func HashFileName(fileName string) string {
	// 获取原始文件的拓展名
	extension := filepath.Ext(fileName)

	// 创建SHA256哈希对象
	hash := sha256.New()

	// 将文件名（不包括拓展名）和当前时间戳转换为字节数组并进行哈希计算
	timestamp := time.Now().UnixNano()
	hash.Write([]byte(fileName[:len(fileName)-len(extension)]))
	hash.Write([]byte(strconv.FormatInt(timestamp, 10)))

	// 获取哈希值并转换为十六进制字符串
	hashedBytes := hash.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)

	// 构建新的文件名，添加原始拓展名
	newFileName := hashedString + "_" + strconv.FormatInt(timestamp, 10) + extension

	return newFileName
}

func GetBaseNameWithoutExtension(fileName string) string {
	extension := filepath.Ext(fileName)
	return strings.TrimSuffix(fileName, extension)
}
