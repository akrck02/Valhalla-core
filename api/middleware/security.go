package middleware

import (
	"context"

	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/error"
	"github.com/akrck02/valhalla-core/log"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const AUTHORITATION_HEADER = "Authorization"

// Manage security
//
// [return] gin.HandlerFunc: handler
func Security(endpoints []models.Endpoint, baseUrl string) gin.HandlerFunc {
	return func(c *gin.Context) {

		var isRegistered = false

		// Check if endpoint is registered and secured
		for _, endpoint := range endpoints {
			if baseUrl+endpoint.Path == c.Request.URL.Path {
				if !endpoint.Secured {
					log.FormattedInfo("Endpoint ${0} is not secured", endpoint.Path)
					return
				}

				isRegistered = true
			}
		}

		// Check if endpoint is registered
		if !isRegistered {
			log.FormattedInfo("Endpoint ${0} is not registered", c.Request.URL.Path)
			return
		}
		log.FormattedInfo("Endpoint ${0} is secured", c.Request.URL.Path)

		// Get token
		token := c.Request.Header.Get(AUTHORITATION_HEADER)

		// Check if token is valid
		if token == "" {
			c.AbortWithStatusJSON(
				utils.HTTP_STATUS_FORBIDDEN,
				gin.H{"code": error.INVALID_TOKEN, "message": "Missing token"},
			)
			return
		}

		// Create a database connection
		var client = db.CreateClient()
		var conn = db.Connect(*client)
		defer db.Disconnect(*client, conn)

		// Check if token is valid
		user, err := IsTokenValid(client, token)

		if err != nil {
			c.AbortWithStatusJSON(
				err.Code,
				gin.H{"code": err.Error, "message": err.Message},
			)

			return
		}

		// Get request

		var request, _ = c.Get("request")
		request = request.(models.Request)

		var castedRequest = request.(models.Request)
		castedRequest.User = user

		// Set user in request
		c.Set("request", castedRequest)
	}
}

// Get user from token
//
//	[param] conn | context.Context : The connection to the database
//	[param] client | *mongo.Client : The client to the database
//	[param] token | *string : The token to check
//	[param] tokenUser | *models.User : The user found or empty --> *models.Error: error if any
func getUserFromToken(conn context.Context, client *mongo.Client, token string) (models.User, *models.Error) {

	var tokenDevice models.Device

	devices := client.Database(db.CurrentDatabase).Collection(db.DEVICE)
	err := devices.FindOne(conn, bson.M{"token": token}).Decode(&tokenDevice)

	if err != nil {
		return models.User{}, &models.Error{
			Code:    utils.HTTP_STATUS_FORBIDDEN,
			Error:   int(error.INVALID_TOKEN),
			Message: "User not matching token",
		}
	}

	var tokenUser models.User

	users := client.Database(db.CurrentDatabase).Collection(db.USER)
	err = users.FindOne(conn, bson.M{"email": tokenDevice.User}).Decode(&tokenUser)

	if err != nil {
		return models.User{}, &models.Error{
			Code:    utils.HTTP_STATUS_FORBIDDEN,
			Error:   int(error.INVALID_TOKEN),
			Message: "User not matching token",
		}
	}

	return tokenUser, nil
}

// Get  if token is valid
//
//	[param] token | string : The token to check
//
//	[return] bool : True if token is valid --> *models.Error: error if any
func IsTokenValid(client *mongo.Client, token string) (*models.User, *models.Error) {

	// decode token
	claims, err := utils.DecryptToken(token)

	if err != nil {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_FORBIDDEN,
			Error:   int(error.INVALID_TOKEN),
			Message: "invalid token format",
		}
	}

	// log token claims
	log.Info("device: " + claims.Claims.(jwt.MapClaims)["device"].(string))
	log.Info("username: " + claims.Claims.(jwt.MapClaims)["username"].(string))
	log.Info("email: " + claims.Claims.(jwt.MapClaims)["email"].(string))

	email := claims.Claims.(jwt.MapClaims)["email"].(string)

	foundUser, tokenUserErr := getUserFromToken(context.Background(), client, token)

	if tokenUserErr != nil {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_FORBIDDEN,
			Error:   int(error.INVALID_TOKEN),
			Message: "invalid token",
		}
	}

	if foundUser.Email != email {
		return nil, &models.Error{
			Code:    utils.HTTP_STATUS_FORBIDDEN,
			Error:   int(error.INVALID_TOKEN),
			Message: "invalid token",
		}
	}

	return &foundUser, nil
}
