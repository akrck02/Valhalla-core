package services

import (
	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

// Create team HTTP API endpoint
//
// [param] c | *gin.Context: gin context
func CreateTeamHttp(c *gin.Context) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params models.Team
	err1 := utils.ReadBodyJson(c, &params)

	if err1 != nil {
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

	var err2 = CreateTeam(conn, client, team)
	if err2 != nil {
		utils.SendResponse(c,
			err2.Code,
			gin.H{"http-code": err2.Code, "internal-code": err2.err2, "message": err2.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "User created"},
	)
}

// Edit team HTTP API endpoint
//
// [param] c | *gin.Context: gin context
func EditTeamHttp(c *gin.Context) {
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params models.Team
	err1 := utils.ReadBodyJson(c, &params)

	if err1 != nil {
		utils.SendResponse(c,
			utils.HTTP_STATUS_BAD_REQUEST,
			gin.H{"code": utils.HTTP_STATUS_NOT_ACCEPTABLE, "message": "Invalid request"},
		)
		return
	}

	var err2 = EditTeam(conn, client, params)

	if err2 != nil {
		utils.SendResponse(c,
			err2.Code,
			gin.H{"http-code": err2.Code, "internal-code": err2.err2, "message": err2.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "Team changed"},
	)
}

// Edit team owner HTTP API endpoint
//
// [param] c | *gin.Context: gin context
func EditTeamOwnerHttp(c *gin.Context) {
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params models.Team
	err1 := utils.ReadBodyJson(c, &params)

	if err1 != nil {
		utils.SendResponse(c,
			utils.HTTP_STATUS_BAD_REQUEST,
			gin.H{"code": utils.HTTP_STATUS_NOT_ACCEPTABLE, "message": "Invalid request"},
		)
		return
	}

	var err2 = EditTeamOwner(conn, client, params)

	if err2 != nil {
		utils.SendResponse(c,
			err2.Code,
			gin.H{"http-code": err2.Code, "internal-code": err2.err2, "message": err2.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "Team owner edited"},
	)
}

// Delete team HTTP API endpoint
//
// [param] c | *gin.Context: gin context
func DeleteTeamHttp(c *gin.Context) {
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params models.Team
	err1 := utils.ReadBodyJson(c, &params)

	if err1 != nil {
		utils.SendResponse(c,
			utils.HTTP_STATUS_BAD_REQUEST,
			gin.H{"code": utils.HTTP_STATUS_NOT_ACCEPTABLE, "message": "Invalid request"},
		)
		return
	}

	err2 := DeleteTeam(conn, client, params)

	if err2 != nil {
		utils.SendResponse(c,
			err2.Code,
			gin.H{"http-code": err2.Code, "internal-code": err1.err2, "message": err2.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "Team deleted"},
	)
}

// Get team HTTP API endpoint
//
// [param] c | *gin.Context: gin context
func GetTeamHttp(c *gin.Context) {
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)
	var params models.Team

	err1 := utils.ReadBodyJson(c, &params)

	if err1 != nil {
		utils.SendResponse(c,
			utils.HTTP_STATUS_BAD_REQUEST,
			gin.H{"code": utils.HTTP_STATUS_NOT_ACCEPTABLE, "message": "Invalid request"},
		)
		return
	}

	team, err2 := GetTeam(conn, client, params)

	if err2 != nil {
		utils.SendResponse(c,
			err2.Code,
			gin.H{"http-code": err2.Code, "internal-code": err1.err2, "message": err2.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "Team found", "team": team},
	)

}

// Add user to team HTTP API endpoint
//
// [param] c | *gin.Context: gin context
func AddMemberHTTP(c *gin.Context) {
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)
	var params MemberChangeRequest

	err1 := utils.ReadBodyJson(c, &params)

	if err1 != nil {
		utils.SendResponse(c,
			utils.HTTP_STATUS_BAD_REQUEST,
			gin.H{"code": utils.HTTP_STATUS_NOT_ACCEPTABLE, "message": "Invalid request"},
		)
		return
	}

	err2 := AddMember(conn, client, params)

	if err2 != nil {
		utils.SendResponse(c,
			err2.Code,
			gin.H{"http-code": err2.Code, "internal-code": err1.err2, "message": err2.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "Member added"},
	)
}
