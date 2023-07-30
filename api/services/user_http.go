package services

import (
	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/log"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

// Register HTTP API endpoint
//
// [param] request | models.Request: request
func RegisterHttp(request models.Request) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user *models.User = &models.User{
		Username: request.GetParamString("username"),
		Password: request.GetParamString("password"),
		Email:    request.GetParamString("email"),
	}

	var error = Register(conn, client, user)
	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: gin.H{"message": "User created"},
	}, nil
}

// Login HTTP API endpoint
//
// [param] request | models.Request: request
func LoginHttp(request models.Request) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user *models.User = &models.User{
		Email:    request.GetParamString("email"),
		Password: request.GetParamString("password"),
	}

	ip := request.IP
	address := request.UserAgent
	token, error := Login(conn, client, user, ip, address)

	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: gin.H{"message": "User found", "auth": token},
	}, nil
}

// Edit user HTTP API endpoint
//
// [param] request | models.Request: request
func EditUserHttp(request models.Request) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user *models.User = &models.User{
		Email:      request.GetParamString("email"),
		Username:   request.GetParamString("username"),
		Password:   request.GetParamString("password"),
		ProfilePic: request.GetParamString("profilePic"),
	}

	// Get token from header
	var token = request.Authorization

	// Get user from token
	tokenUser, error := GetUserFromToken(conn, client, token)
	if error != nil {
		return nil, error
	}

	// Check if user is the same
	if tokenUser.Email != user.Email {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   utils.HTTP_STATUS_NOT_ACCEPTABLE,
			Message: "Invalid request",
		}
	}

	updateErr := EditUser(conn, client, user)
	if updateErr != nil {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   utils.HTTP_STATUS_NOT_ACCEPTABLE,
			Message: "Invalid request",
		}
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: gin.H{"message": "User updated"},
	}, nil
}

// Change email HTTP API endpoint
//
// [param] request | models.Request: request
func EditUserEmailHttp(request models.Request) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var email *EmailChangeRequest = &EmailChangeRequest{
		Email:    request.GetParamString("email"),
		NewEmail: request.GetParamString("newEmail"),
	}

	changeErr := EditUserEmail(conn, client, email)
	if changeErr != nil {
		return nil, changeErr
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: gin.H{"message": "Email changed"},
	}, nil
}

// Delete user HTTP API endpoint
//
// [param] request | models.Request: request
func DeleteUserHttp(request models.Request) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user *models.User = &models.User{
		Email: request.GetParamString("email"),
	}

	deleteErr := DeleteUser(conn, client, user)
	if deleteErr != nil {
		return nil, deleteErr
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: gin.H{"message": "User deleted"},
	}, nil
}

// Get user HTTP API endpoint
//
// [param] request | models.Request: request
func GetUserHttp(request models.Request) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	// Get code from url GET parameter
	id := request.GetParamString("id")
	if id == "" {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   utils.HTTP_STATUS_NOT_ACCEPTABLE,
			Message: "Id cannot be empty",
		}
	}

	var user *models.User = &models.User{
		Email: id,
	}

	var foundUser, error = GetUser(conn, client, user, true)
	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: gin.H{"message": "User found", "user": foundUser},
	}, nil

}

// Edit user profile picture HTTP API endpoint
//
// [param] request | models.Request: request
func EditUserProfilePictureHttp(request models.Request) (*models.Response, *models.Error) {

	// Get user
	var user *models.User = &models.User{
		Email: request.GetParamString("email"),
	}

	// Get image
	bytes := request.Body
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	// Upload image
	var error = EditUserProfilePicture(conn, client, user, bytes)
	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: gin.H{"message": "Profile picture updated"},
	}, nil

}

// Validate user account HTTP API endpoint
//
// [param] request | models.Request: request
func ValidateUserHttp(request models.Request) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	// Get code from url GET parameter
	code := request.GetParamString("code")
	log.Info("Query code: " + code)

	var error = ValidateUser(conn, client, code)
	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     utils.HTTP_STATUS_OK,
		Response: gin.H{"message": "User validated"},
	}, nil
}

// Get
