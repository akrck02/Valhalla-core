package services

import (
	"github.com/akrck02/valhalla-core/error"
)

type validatorResponse struct {
	Response error.Validator
	Message  string
}

func isOwnerOfTeam(token string, _id string) validatorResponse {

	// Check if token is associated with user with id identical to owner parameter in team

	return validatorResponse{
		Response: 200,
		Message:  "Ok.",
	}
}
