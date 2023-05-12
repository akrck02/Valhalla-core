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
	team.Description = params.Description
	team.Owner = params.Owner
	team.ProfilePic = params.ProfilePic

	var error = CreateTeam(conn, client, team)
	if error != nil {
		utils.SendResponse(c,
			error.Code,
			gin.H{"http-code": error.Code, "internal-code": error.Error, "message": error.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "User created"},
	)
}

func EditTeamHttp(c *gin.Context) {
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

	var error = EditTeam(conn, client, params)

	if error != nil {
		utils.SendResponse(c,
			error.Code,
			gin.H{"http-code": error.Code, "internal-code": error.Error, "message": error.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "Team changed"},
	)
}

func EditTeamOwnerHttp(c *gin.Context) {
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

	var error = EditTeamOwner(conn, client, params)

	if error != nil {
		utils.SendResponse(c,
			error.Code,
			gin.H{"http-code": error.Code, "internal-code": error.Error, "message": error.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "Team owner edited"},
	)
}

func DeleteTeamHttp(c *gin.Context) {
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

	error := DeleteTeam(conn, client, params)

	if error != nil {
		utils.SendResponse(c,
			error.Code,
			gin.H{"http-code": error.Code, "internal-code": err.Error, "message": error.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "Team deleted"},
	)
}

func GetTeamHttp(c *gin.Context) {
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

	team, error := GetTeam(conn, client, params)

	if error != nil {
		utils.SendResponse(c,
			error.Code,
			gin.H{"http-code": error.Code, "internal-code": err.Error, "message": error.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "Team found", "team": team},
	)

}
