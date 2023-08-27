package utils

import "strings"

/*
IsVideoFile 检查文件是否为视频文件
params: fileName 文件名
return: bool 是true 非false
*/
func IsVideoFile(fileName string) bool {
	videoExtensions := []string{".mp4", ".avi", ".mkv", ".mov", ".wmv", ".flv", ".webm", ".m4v"}

	ext := strings.ToLower(fileName[strings.LastIndex(fileName, "."):])
	for _, videoExt := range videoExtensions {
		if ext == videoExt {
			return true
		}
	}
	return false
}
