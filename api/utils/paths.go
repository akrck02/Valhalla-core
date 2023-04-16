package utils

import "github.com/akrck02/valhalla-core/configuration"

func GetProfilePicturePath(username string) string {
	return configuration.PROFILE_PICTURES_PATH + username + ".jpg"
}
