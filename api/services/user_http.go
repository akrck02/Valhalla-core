package services

import (
	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

// Register HTTP API endpoint
//
// [param] c | *gin.Context: gin context
func RegisterHttp(c *gin.Context) {

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

	var user models.User
	user.Username = params.Username
	user.Password = params.Password
	user.Email = params.Email

	var error = Register(conn, client, user)
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

// Login HTTP API endpoint
//
// [param] c | *gin.Context: gin context
func LoginHttp(c *gin.Context) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user models.User
	err := utils.ReadBodyJson(c, &user)

	if err != nil {
		utils.SendResponse(c,
			utils.HTTP_STATUS_BAD_REQUEST,
			gin.H{"code": utils.HTTP_STATUS_NOT_ACCEPTABLE, "message": "Invalid request"},
		)
		return
	}

	ip := c.ClientIP()
	address := c.Request.Header.Get("User-Agent")
	token, error := Login(conn, client, user, ip, address)

	if error != nil {
		utils.SendResponse(c,
			error.Code,
			gin.H{"http-code": error.Code, "internal-code": error.Error, "message": error.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"code": utils.HTTP_STATUS_OK, "message": "User found", "auth": token},
	)
}

// Edit user HTTP API endpoint
//
// [param] c | *gin.Context: gin context
func EditUserHttp(c *gin.Context) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user models.User
	err := utils.ReadBodyJson(c, &user)

	if err != nil {
		utils.SendResponse(c,
			utils.HTTP_STATUS_BAD_REQUEST,
			gin.H{"code": utils.HTTP_STATUS_NOT_ACCEPTABLE, "message": "Invalid request"},
		)
		return
	}

	updateErr := EditUser(conn, client, user)
	if updateErr != nil {
		utils.SendResponse(c,
			updateErr.Code,
			gin.H{"http-code": updateErr.Code, "internal-code": updateErr.Error, "message": updateErr.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "User updated"},
	)
}

// Change email HTTP API endpoint
//
// [param] c | *gin.Context: gin context
// [return] *models.Error: error if any
func EditUserEmailHttp(c *gin.Context) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var email EmailChangeRequest
	err := utils.ReadBodyJson(c, &email)

	if err != nil {
		utils.SendResponse(c,
			utils.HTTP_STATUS_BAD_REQUEST,
			gin.H{"code": utils.HTTP_STATUS_NOT_ACCEPTABLE, "message": "Invalid request"},
		)
		return
	}

	changeErr := EditUserEmail(conn, client, email)
	if changeErr != nil {
		utils.SendResponse(c,
			changeErr.Code,
			gin.H{"http-code": changeErr.Code, "internal-code": changeErr.Error, "message": changeErr.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "Email changed"},
	)
}

// Delete user HTTP API endpoint
//
// [param] c | *gin.Context: gin context
// [return] *models.Error: error if any
func DeleteUserHttp(c *gin.Context) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user models.User
	err := utils.ReadBodyJson(c, &user)

	if err != nil {
		utils.SendResponse(c,
			utils.HTTP_STATUS_BAD_REQUEST,
			gin.H{"code": utils.HTTP_STATUS_NOT_ACCEPTABLE, "message": "Invalid request"},
		)
		return
	}

	deleteErr := DeleteUser(conn, client, user)
	if deleteErr != nil {
		utils.SendResponse(c,
			deleteErr.Code,
			gin.H{"http-code": deleteErr.Code, "internal-code": deleteErr.Error, "message": deleteErr.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "User deleted"},
	)
}

// Get user HTTP API endpoint
func GetUserHttp(c *gin.Context) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var user models.User
	err := utils.ReadBodyJson(c, &user)

	if err != nil {
		utils.SendResponse(c,
			utils.HTTP_STATUS_BAD_REQUEST,
			gin.H{"code": utils.HTTP_STATUS_NOT_ACCEPTABLE, "message": "Invalid request"},
		)
		return
	}

	var foundUser models.User
	var error = GetUser(conn, client, user, &foundUser)
	if error != nil {
		utils.SendResponse(c,
			error.Code,
			gin.H{"http-code": error.Code, "internal-code": error.Error, "message": error.Message},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"http-code": utils.HTTP_STATUS_OK, "message": "User found", "user": foundUser},
	)

}
