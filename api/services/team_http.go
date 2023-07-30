package services

import (
	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

// Create team HTTP API endpoint
//
// [param] request | models.Request: request
func CreateTeamHttp(request models.Request) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var team *models.Team
	team = team.FromRequest(request)

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
// [param] request | models.Request: request
func EditTeamHttp(request models.Request) (*models.Response, *models.Error) {
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params *models.Team
	params = params.FromRequest(request)

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
// [param] request | models.Request: request
func EditTeamOwnerHttp(request models.Request) (*models.Response, *models.Error) {
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params *models.Team
	params = params.FromRequest(request)

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
// [param] request | models.Request: request
func DeleteTeamHttp(request models.Request) (*models.Response, *models.Error) {
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params *models.Team
	params = params.FromRequest(request)

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
func GetTeamHttp(request models.Request) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params models.Team = models.Team{}
	params.ID = request.GetParamString("id")

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
// [param] request | models.Request: request
func AddMemberHttp(request models.Request) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)
	var params *MemberChangeRequest
	params = params.FromRequest(request)

	var err = AddMember(conn, client, params)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: gin.H{"message": "Member added"},
	}, nil
}

// Remove user from team HTTP API endpoint
//
// [param] request | models.Request: request
func RemoveMemberHttp(request models.Request) (*models.Response, *models.Error) {
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params *MemberChangeRequest
	params = params.FromRequest(request)

	var err = RemoveMember(conn, client, params)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: gin.H{"message": "Member removed"},
	}, nil
}
