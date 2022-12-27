package util

import "time"

// TimeNow 获取当前时间
func TimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
