package services

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/error"
	"github.com/akrck02/valhalla-core/lang"
	"github.com/akrck02/valhalla-core/log"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RegisterParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type validatePasswordResult struct {
	Response error.User
	Message  string
}

const MINIMUM_CHARACTERS_FOR_PASSWORD = 16

var SPECIAL_CHARATERS = []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "=", "+", "[", "]", "{", "}", "|", ";", ":", "'", ",", ".", "<", ">", "?", "/", "`", "~"}

func Register(c *gin.Context) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	var params RegisterParams
	err := c.ShouldBindJSON(&params)

	if err != nil {
		utils.SendResponse(c,
			utils.HTTP_STATUS_NOT_ACCEPTABLE,
			gin.H{"code": utils.HTTP_STATUS_NOT_ACCEPTABLE, "message": "Invalid request"},
		)
		return
	}

	var user models.User
	user.Username = params.Username
	user.Password = params.Password
	user.Email = params.Email

	var checkedPass = validatePassword(user.Password)

	if checkedPass.Response != 200 {
		utils.SendResponse(c,
			utils.HTTP_STATUS_FORBIDDEN,
			gin.H{"code": utils.HTTP_STATUS_FORBIDDEN, "error": checkedPass.Response, "message": checkedPass.Message},
		)
		return
	}

	coll := client.Database("valhalla").Collection("users")
	found := mailExists(user.Email, conn, coll)

	if found.Email != "" {
		utils.SendResponse(c,
			utils.HTTP_STATUS_CONFLICT,
			gin.H{"code": utils.HTTP_STATUS_CONFLICT, "message": "User already exists"},
		)
		return
	}

	// register user on database
	usr, err := coll.InsertOne(conn, user)

	if err != nil {
		utils.SendResponse(c,
			utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			gin.H{"code": utils.HTTP_STATUS_INTERNAL_SERVER_ERROR, "message": err.Error()},
		)
		return
	}

	// send response
	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"code": utils.HTTP_STATUS_OK, "message": "User created", "data": usr},
	)
}

// Check the given credentials.
//
//	[HTTP]  POST
//	[param] username | string: username of the user
//	[returns] the user can login or not
func Login(c *gin.Context) {

	var client = db.CreateClient()
	var conn = db.Connect(*client)
	defer db.Disconnect(*client, conn)

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var user models.User
	json.Unmarshal([]byte(jsonData), &user)

	coll := client.Database("valhalla").Collection("users")
	found := authorizationOk(user.Username, user.Password, conn, coll)

	if found.Email == "" {
		utils.SendResponse(c,
			utils.HTTP_STATUS_FORBIDDEN,
			gin.H{"code": utils.HTTP_STATUS_NOT_FOUND, "message": "Forbidden"},
		)
		return
	}

	utils.SendResponse(c,
		utils.HTTP_STATUS_OK,
		gin.H{"code": utils.HTTP_STATUS_OK, "message": "User found", "data": found},
	)
	log.Info(user.Username + " / " + user.Password)
}

// Check if the given password is valid
// following the next rules:
//
//		[-] At least 16 characters
//		[-] At least one special character
//		[-] At least one number
//
//	 [param] password | string: password to check
//	 [returns] the password is valid or not
func validatePassword(password string) validatePasswordResult {

	if len(password) < MINIMUM_CHARACTERS_FOR_PASSWORD {
		return validatePasswordResult{
			Response: error.SHORT_PASSWORD,
			Message:  "Password must have at least " + lang.Int2String(MINIMUM_CHARACTERS_FOR_PASSWORD) + " characters",
		}
	}

	if !utils.ContainsAny(password, SPECIAL_CHARATERS) {
		return validatePasswordResult{
			Response: error.NO_SPECIAL_CHARACTERS_PASSWORD,
			Message:  "Password must have at least one special character",
		}
	}

	return validatePasswordResult{
		Response: 200,
		Message:  "Ok.",
	}
}

// Check email on database
//
//	[param] email | string The email to check
//	[param] conn | context.Context The connection to the database
//	[returns] model.User | The user found or empty
func mailExists(email string, conn context.Context, coll *mongo.Collection) models.User {

	filter := bson.D{{Key: "email", Value: email}}

	var result models.User
	coll.FindOne(conn, filter).Decode(&result)

	return result

}

// Get if the given credentials are valid
//
//	[param] username | string : The username to check
//	[param] password | string : The password to check
//	[param] conn | context.Context : The connection to the database
//	[returns] model.User | The user found or empty
func authorizationOk(username string, password string, conn context.Context, coll *mongo.Collection) models.User {

	filter := bson.D{{Key: "username", Value: username}, {Key: "password", Value: password}}

	var result models.User
	coll.FindOne(conn, filter).Decode(&result)

	return result

}
