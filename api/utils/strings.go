package utils

import (
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ContainsAny returns true if the string contains any of the chars
//
// [param] str | string: string to check
// [param] chars | []string: chars to check
//
// [return] bool: true if the string contains any of the chars
func ContainsAny(str string, chars []string) bool {
	return strings.ContainsAny(str, strings.Join(chars, ""))
}

// IsLowerCase returns true if the string is lowercase
//
// [param] str | string: string to check
//
// [return] bool: true if the string is lowercase
func IsLowerCase(str string) bool {
	return str == strings.ToLower(str)
}

// IsUpperCase returns true if the string is uppercase
//
// [param] str | string: string to check
//
// [return] bool: true if the string is uppercase
func IsUpperCase(str string) bool {
	return str == strings.ToUpper(str)
}

// ContainsNumbers returns true if the string contains any number
//
// [param] str | string: string to check
//
// [return] bool: true if the string contains any number
func ContainsNumbers(str string) bool {
	return strings.ContainsAny(str, "0123456789")
}

// ContainsSpecialCharacters returns true if the string contains any special character
//
// [param] str | string: string to check
//
// [return] bool: true if the string contains any special character
func ContainsSpecialCharacters(str string) bool {
	return strings.ContainsAny(str, "!@#$%^&*()_+{}|:\"<>?,./;'[]\\-=")
}

// IsEmpty returns true if the string is empty
//
// [param] str | string: string to check
//
// [return] bool: true if the string is empty
func IsEmpty(str string) bool {
	return str == ""
}

// StringToObjectId converts a string to a primitive.ObjectID
//
// [param] str | string: string to convert
//
// [return] primitive.ObjectID: converted string
func StringToObjectId(str string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(str)
}

// Convert Int64 to String
//
// [param] num | int64: number to convert
//
// [return] string: converted number
func Int642String(num int64) string {
	return strconv.FormatInt(num, 10)
}

// Convert int to string
//
// [param] num | int: number to convert
//
// [return] string: converted number
func Int2String(num int) string {
	return strconv.Itoa(num)
}

// Get string pointer from string or nil if empty
//
// [param] str | string: string to convert
//
// [return] *string: string pointer
func ValueOrEmpty(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}
