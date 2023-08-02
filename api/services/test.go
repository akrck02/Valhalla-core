package services

import (
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

type pingResponse struct {
	Message string `json:"message"`
}

// Ping HTTP API endpoint
//
// [param] c | *gin.Context: context
//
// [return] *models.Response: response	| *models.Error: error
func PingHttp(c *gin.Context) (*models.Response, *models.Error) {

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: pingResponse{Message: "pong"},
	}, nil
}
