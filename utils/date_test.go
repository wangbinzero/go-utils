package utils

import (
	"testing"
	"time"
)

func TestUTCFormatString(t *testing.T) {
	now := time.Now().Unix()
	val := UnixToString(now, "2006-01-02 15:04:05")
	t.Log("UTC时间转换字符串:", val)
}

func TestISO8601FormatString(t *testing.T) {
	now := time.Now().Unix()
	val := ISO8601FormatString(now)
	t.Log("UTC时间转换为ISO8601:", val)
}

func TestStringToUnix(t *testing.T) {
	time := "2019-09-10 12:09:08"
	val := StringToUnix(time, "2006-01-02 15:04:05")
	t.Log("字符串转unix:", val)
}
