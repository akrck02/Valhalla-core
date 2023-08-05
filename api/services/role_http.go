package services

import (
	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/error"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

// CreateRole HTTP API endpoint
//
// [param] c | *gin.Context: context
//
// [return] *models.Response: response | *models.Error: error
func CreateRoleHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params models.User
	err := c.ShouldBindJSON(&params)

	if err != nil {
		return nil, &models.Error{
			Status:  utils.HTTP_STATUS_NOT_ACCEPTABLE,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	var role models.Role
	var error = CreateRole(conn, client, role)

	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: "User created",
	}, nil

}

// DeleteRole HTTP API endpoint
//
// [param] c | *gin.Context: context
//
// [return] *models.Response: response | *models.Error: error
func DeleteRoleHttp(c *gin.Context) (*models.Response, *models.Error) {
	return nil, &models.Error{
		Status:  utils.HTTP_STATUS_NOT_IMPLEMENTED,
		Error:   error.NOT_IMPLEMENTED,
		Message: "Not implemented",
	}
}

// EditRole HTTP API endpoint
//
// [param] c | *gin.Context: context
//
// [return] *models.Response: response | *models.Error: error
func EditRoleHttp(c *gin.Context) (*models.Response, *models.Error) {
	return nil, &models.Error{
		Status:  utils.HTTP_STATUS_NOT_IMPLEMENTED,
		Error:   error.NOT_IMPLEMENTED,
		Message: "Not implemented",
	}
}

// GetRole HTTP API endpoint
//
// [param] c | *gin.Context: context
//
// [return] *models.Response: response | *models.Error: error
func GetRoleHttp(c *gin.Context) (*models.Response, *models.Error) {
	return nil, &models.Error{
		Status:  utils.HTTP_STATUS_NOT_IMPLEMENTED,
		Error:   error.NOT_IMPLEMENTED,
		Message: "Not implemented",
	}
}
