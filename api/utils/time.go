package utils

import "time"

func getCurrentMillis() int64 {
	return time.Now().UnixMilli()
}
