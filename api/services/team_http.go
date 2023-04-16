package services

import (
	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

func CreateTeamHttp(c *gin.Context) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params models.Team
	err := utils.ReadBodyJson(c, &params)

	if err != nil {
		utils.SendResponse(c,
			utils.HTTP_STATUS_BAD_REQUEST,
			gin.H{"code": utils.HTTP_STATUS_NOT_ACCEPTABLE, "message": "Invalid request"},
		)
		return
	}

	var team models.Team

	team.Name = params.Name

}

func EditTeam(c *gin.Context) {
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params models.Team
	err := utils.ReadBodyJson(c, &params)

	if err != nil {
		utils.SendResponse(c,
			utils.HTTP_STATUS_BAD_REQUEST,
			gin.H{"code": utils.HTTP_STATUS_NOT_ACCEPTABLE, "message": "Invalid request"},
		)
		return
	}

	var team models.Team

	team.Name = params.Name
}

func DeleteTeam(c *gin.Context) {

}
