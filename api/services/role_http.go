package services

import (
	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

func CreateRoleHttp(c *gin.Context) {
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params models.User
	err := c.ShouldBindJSON(&params)

	if err != nil {
		utils.SendResponse(c,
			utils.HTTP_STATUS_BAD_REQUEST,
			gin.H{"code": utils.HTTP_STATUS_NOT_ACCEPTABLE, "message": "Invalid request"},
		)
		return
	}

	var role models.Role

	var error = CreateRole(conn, client, role)

	if error != nil {
		utils.SendResponse(c,
			error.Code,
			gin.H{"http-code": error.Code, "internal-code": error.Error, "message": error.Message},
		)
		return
	}

	// send response
	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "User created"},
	)

}

func DeleteRoleHttp(c *gin.Context) {

}

func EditRoleHttp(c *gin.Context) {

}

func GetRoleHttp(c *gin.Context) {

}
