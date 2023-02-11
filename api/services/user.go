package services

import (
	"context"
	"net/http"

	"github.com/akrck02/valhalla-core/models"
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

func Register(c *gin.Context, logger *log.Logger, conn context.Context, client mongo.Client) {

	var params RegisterParams
	err := c.ShouldBindJSON(&params)

	if err != nil {

		//406
		SendResponse(c, helpers.Response{Status: http.StatusUnauthorized, Error: []string{"Username and password do not match"}})
		logger.Error(err)
	}

	var user models.User
	user.Username = params.Username
	user.Password = params.Password
	user.Email = params.Email

	if validatePassword(user.Password) == INVALID_PASSWORD {
		// return message with reasoning for invalid value

		return
	}

	coll := client.Database("Valhalla").Collection("users")
	found := exists("", conn, coll)

	logger.Info(found)
}

func Login(c *gin.Context) {

}

func validatePassword(password string) validatePasswordResponse {

	return VALID_PASSWORD
}

func exists(email string, conn context.Context, coll *mongo.Collection) models.User {

	filter := bson.D{{"email", email}}

	var result models.User
	coll.FindOne(conn, filter).Decode(&result)

	return result

}
