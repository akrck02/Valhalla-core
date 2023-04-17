package mock

import (
	"github.com/akrck02/valhalla-core/configuration"
	"github.com/akrck02/valhalla-core/utils"
)

func ProfilePicture() ([]byte, error) {

	profilePic, readErr := utils.ReadFile(configuration.BASE_PATH + "resources/images/cat.jpg")
	return profilePic, readErr

}
