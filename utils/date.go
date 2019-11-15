package utils

import "time"

// int64格式转换为字符串
func UnixToString(val int64, format string) string {
	return time.Unix(val, 0).Format(format)
}

// int64格式转换为字符串
func ISO8601FormatString(val int64) string {
	return time.Unix(val, 0).Format(time.RFC3339)
}

// 字符串转int64
func StringToUnix(val, template string) int64 {
	unix, _ := time.ParseInLocation(template, val, time.Local)
	return unix.Unix()
}
