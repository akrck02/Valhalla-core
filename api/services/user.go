package services

import (
	"context"

	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
	"github.com/withmandala/go-log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RegisterParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type validatePasswordResponse int64

const (
	VALID_PASSWORD   = 0
	INVALID_PASSWORD = 1
)

func Register(c *gin.Context, logger *log.Logger) {

	var client = db.CreateClient(logger)
	var conn = db.Connect(logger, *client)
	defer db.Disconnect(logger, *client, conn)

	var params RegisterParams
	err := c.ShouldBindJSON(&params)

	if err != nil {
		utils.SendResponse(c,
			utils.HTTP_STATUS_NOT_ACCEPTABLE,
			gin.H{"code": utils.HTTP_STATUS_NOT_ACCEPTABLE, "message": "Invalid request"},
		)
	}

	var user models.User
	user.Username = params.Username
	user.Password = params.Password
	user.Email = params.Email

	if validatePassword(user.Password) == INVALID_PASSWORD {
		utils.SendResponse(c,
			utils.HTTP_STATUS_FORBIDDEN,
			gin.H{"code": utils.HTTP_STATUS_FORBIDDEN, "message": "Invalid password"},
		)
		return
	}

	coll := client.Database("valhalla").Collection("users")
	found := exists(user.Email, conn, coll)

	logger.Info(found)

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

func Login(c *gin.Context) {

}

func validatePassword(password string) validatePasswordResponse {

	return VALID_PASSWORD
}

func exists(email string, conn context.Context, coll *mongo.Collection) models.User {

	filter := bson.D{{Key: "email", Value: email}}

	var result models.User
	coll.FindOne(conn, filter).Decode(&result)

	return result

}
