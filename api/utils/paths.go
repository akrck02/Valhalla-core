package utils

import (
	"os"

	"github.com/akrck02/valhalla-core/configuration"
)

const DEFAULT_PROFILE_PICTURE_EXTENSION = ".jpg"

func GetProfilePicturePath(username string) string {

	var path string = configuration.PROFILE_PICTURES_PATH + username

	if username != "" {
		path += DEFAULT_PROFILE_PICTURE_EXTENSION
	}

	return path
}

func ExistsDir(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
