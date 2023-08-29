package utils

import "time"

func StringToTime(timestampStr string) (time.Time, error) {
	if timestampStr == "" {
		// 如果字符串为空，返回一个零值时间
		return time.Time{}, nil
	}
	// 使用 time.Parse 进行转换
	parsedTime, err := time.Parse(time.DateTime, timestampStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
