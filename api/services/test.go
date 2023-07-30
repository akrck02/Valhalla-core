package services

import (
	"fmt"

	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
)

type pingResponse struct {
	Message string `json:"message"`
}

/**
 * Status check endpoint
 */
func PingHttp(request models.Request) (*models.Response, *models.Error) {

	fmt.Println(request)

	return &models.Response{
			Code:     utils.HTTP_STATUS_OK,
			Response: pingResponse{Message: "pong"},
		}, &models.Error{
			Code:    500,
			Message: "noooooooooo Maideeeer ",
			Error:   230,
		}

}
