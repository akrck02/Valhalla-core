package services

import (
	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/error"
	"github.com/akrck02/valhalla-core/log"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

// Register HTTP API endpoint
//
// [param] c | *gin.Context: context
func RegisterHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user *models.User = &models.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		return nil, &models.Error{
			Code:  utils.HTTP_STATUS_BAD_REQUEST,
			Error: utils.HTTP_STATUS_NOT_ACCEPTABLE,
		}
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
// [param] c | *gin.Context: context
func LoginHttp(c *gin.Context) (*models.Response, *models.Error) {

	var request = utils.GetRequestMetadata(c)
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user *models.User = &models.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		return nil, &models.Error{
			Code:  utils.HTTP_STATUS_BAD_REQUEST,
			Error: utils.HTTP_STATUS_NOT_ACCEPTABLE,
		}
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
// [param] c | *gin.Context: context
func EditUserHttp(c *gin.Context) (*models.Response, *models.Error) {

	var request = utils.GetRequestMetadata(c)
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var userToEdit *models.User = &models.User{}
	err := c.ShouldBindJSON(userToEdit)
	if err != nil {
		return nil, &models.Error{
			Code:  utils.HTTP_STATUS_BAD_REQUEST,
			Error: utils.HTTP_STATUS_NOT_ACCEPTABLE,
		}
	}

	// get if request user can edit the user
	canEdit := CanEditUser(request.User, userToEdit)
	if !canEdit {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_FORBIDDEN,
			Error:   error.ACCESS_DENIED,
			Message: "Cannot edit user",
		}
	}

	updateErr := EditUser(conn, client, userToEdit)
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
// [param] c | *gin.Context: context
func EditUserEmailHttp(c *gin.Context) (*models.Response, *models.Error) {

	var request = utils.GetRequestMetadata(c)
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var email *EmailChangeRequest = &EmailChangeRequest{}
	err := c.ShouldBindJSON(email)
	if err != nil {
		return nil, &models.Error{
			Code:  utils.HTTP_STATUS_BAD_REQUEST,
			Error: utils.HTTP_STATUS_NOT_ACCEPTABLE,
		}
	}

	// get if request user can edit the user
	canEdit := CanEditUser(request.User, &models.User{Email: email.Email})
	if !canEdit {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_FORBIDDEN,
			Error:   error.ACCESS_DENIED,
			Message: "Access denied: Cannot edit user",
		}
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
// [param] c | *gin.Context: context
func DeleteUserHttp(c *gin.Context) (*models.Response, *models.Error) {

	var request = utils.GetRequestMetadata(c)
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user *models.User = &models.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		return nil, &models.Error{
			Code:  utils.HTTP_STATUS_BAD_REQUEST,
			Error: utils.HTTP_STATUS_NOT_ACCEPTABLE,
		}
	}

	// get if request user can delete the user
	canDelete := CanEditUser(request.User, user)
	if !canDelete {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_FORBIDDEN,
			Error:   error.ACCESS_DENIED,
			Message: "Access denied: Cannot delete user",
		}
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
// [param] c | *gin.Context: context
func GetUserHttp(c *gin.Context) (*models.Response, *models.Error) {

	var request = utils.GetRequestMetadata(c)
	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	// Get code from url GET parameter
	id := c.Query("id")
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

	// get if request user can see the user
	canSee := CanSeeUser(request.User, user)
	if !canSee {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_FORBIDDEN,
			Error:   error.ACCESS_DENIED,
			Message: "Access denied: Cannot see the user",
		}
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
// [param] c | *gin.Context: context
func EditUserProfilePictureHttp(c *gin.Context) (*models.Response, *models.Error) {

	var request = utils.GetRequestMetadata(c)

	// Get user
	var user *models.User = &models.User{
		Email: c.Query("email"),
	}

	// Get image as bytes
	bytes, err := utils.MultipartToBytes(c, "ProfilePicture")
	if err != nil {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_BAD_REQUEST,
			Error:   utils.HTTP_STATUS_NOT_ACCEPTABLE,
			Message: "Invalid request",
		}
	}

	// get if request user can delete the user
	canEdit := CanEditUser(request.User, user)
	if !canEdit {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_FORBIDDEN,
			Error:   error.ACCESS_DENIED,
			Message: "Access denied: Cannot edit user",
		}
	}

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
// [param] c | *gin.Context: context
func ValidateUserHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	// Get code from url GET parameter
	code := c.Query("code")
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
