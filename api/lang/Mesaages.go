package lang

import (
	"strconv"
	"strings"
)

const APP_TITLE = "Valhalla Core - chocolate version"

func Format(message string, args ...string) string {
	var i int

	for i = 0; i < len(args); i++ {
		message = strings.Replace(message, "${"+int2String(i)+"}", args[i], -1)
	}

	return message
}

func int2String(num int) string {
	return strconv.FormatInt(int64(num), 10)
}
