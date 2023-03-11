package lang

import (
	"strconv"
	"strings"
)

const APP_TITLE = "Valhalla Core - chocolate version"

func Format(message string, args ...string) string {
	var i int

	for i = 0; i < len(args); i++ {
		message = strings.Replace(message, "${"+Int2String(i)+"}", args[i], -1)
	}

	return message
}

func Int2String(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

func Int642String(num int64) string {
	return strconv.FormatInt(int64(num), 10)
}
