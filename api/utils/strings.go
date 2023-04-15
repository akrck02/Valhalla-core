package utils

import "strings"

func ContainsAny(str string, chars []string) bool {
	for _, c := range chars {
		if strings.Contains(str, c) {
			return true
		}
	}
	return false
}

func IsLowerCase(str string) bool {
	return str == strings.ToLower(str)
}

func IsUpperCase(str string) bool {
	return str == strings.ToUpper(str)
}
