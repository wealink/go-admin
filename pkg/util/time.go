package util

import "time"

func GetCurrentTime() int64 {
	// return time.Now().Format("2006-01-02 15:04:05")
	return time.Now().Unix()
}