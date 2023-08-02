package services

import (
	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

// Create team HTTP API endpoint
//
// [param] c | *gin.Context: context
func CreateTeamHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var team *models.Team = &models.Team{}

	err := c.ShouldBindJSON(team)
	if err != nil {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   utils.HTTP_STATUS_NOT_ACCEPTABLE,
			Message: "Invalid request body",
		}
	}

	var error = CreateTeam(conn, client, team)
	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: gin.H{"message": "Team created"},
	}, nil

}

// Edit team HTTP API endpoint
//
// [param] c | *gin.Context: context
func EditTeamHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params *models.Team = &models.Team{}
	err := c.ShouldBindJSON(params)
	if err != nil {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   utils.HTTP_STATUS_NOT_ACCEPTABLE,
			Message: "Invalid request body",
		}
	}

	var error = EditTeam(conn, client, params)

	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: gin.H{"message": "Team changed"},
	}, nil
}

// Edit team owner HTTP API endpoint
//
// [param] c | *gin.Context: context
func EditTeamOwnerHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params *models.Team = &models.Team{}

	err := c.ShouldBindJSON(params)
	if err != nil {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   utils.HTTP_STATUS_NOT_ACCEPTABLE,
			Message: "Invalid request body",
		}
	}

	var error = EditTeamOwner(conn, client, params)

	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: gin.H{"message": "Team owner edited"},
	}, nil
}

// Delete team HTTP API endpoint
//
// [param] c | *gin.Context: context
func DeleteTeamHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params *models.Team = &models.Team{}
	err := c.ShouldBindJSON(params)
	if err != nil {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   utils.HTTP_STATUS_NOT_ACCEPTABLE,
			Message: "Invalid request body",
		}
	}

	var error = DeleteTeam(conn, client, params)
	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: gin.H{"message": "Team deleted"},
	}, nil
}

// Get team HTTP API endpoint
//
// [param] request | models.Request: request
func GetTeamHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params models.Team = models.Team{}
	params.ID = c.Query("id")

	if params.ID == "" {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   utils.HTTP_STATUS_NOT_ACCEPTABLE,
			Message: "Team ID is required",
		}
	}

	team, error := GetTeam(conn, client, &params)
	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: gin.H{"message": "Team found", "team": team},
	}, nil
}

// Add user to team HTTP API endpoint
//
// [param] c | *gin.Context: context
func AddMemberHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)
	var params *MemberChangeRequest = &MemberChangeRequest{}

	err := c.ShouldBindJSON(params)
	if err != nil {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   utils.HTTP_STATUS_NOT_ACCEPTABLE,
			Message: "Invalid request body",
		}
	}

	var addMemberErr = AddMember(conn, client, params)
	if err != nil {
		return nil, addMemberErr
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: gin.H{"message": "Member added"},
	}, nil
}

// Remove user from team HTTP API endpoint
//
// [param] c | *gin.Context: context
func RemoveMemberHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params *MemberChangeRequest = &MemberChangeRequest{}
	err := c.ShouldBindJSON(params)
	if err != nil {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   utils.HTTP_STATUS_NOT_ACCEPTABLE,
			Message: "Invalid request body",
		}
	}

	var removeMemberErr = RemoveMember(conn, client, params)
	if err != nil {
		return nil, removeMemberErr
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: gin.H{"message": "Member removed"},
	}, nil
}
