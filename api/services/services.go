package services

import (
	"github.com/akrck02/valhalla-core/configuration"
	"github.com/akrck02/valhalla-core/db"
	"github.com/akrck02/valhalla-core/error"
	"github.com/akrck02/valhalla-core/log"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

const API_PATH = "api"
const VERSION = "v1"
const API_COMPLETE = "/" + API_PATH + "/" + VERSION + "/"

type Endpoint struct {
	Path     string               `json:"path"`
	Method   int                  `json:"method"`
	Listener func(c *gin.Context) `json:"listener"`
	Secured  bool                 `json:"secured"`
}

var endpoints = []Endpoint{

	// User endpoints
	{"user/register", utils.HTTP_METHOD_POST, RegisterHttp, false},
	{"user/login", utils.HTTP_METHOD_POST, LoginHttp, false},
	{"user/edit", utils.HTTP_METHOD_POST, EditUserHttp, true},
	{"user/edit/email", utils.HTTP_METHOD_POST, EditUserEmailHttp, true},
	{"user/edit/profilepicture", utils.HTTP_METHOD_POST, EditUserProfilePictureHttp, true},
	{"user/delete", utils.HTTP_METHOD_POST, DeleteUserHttp, true},
	{"user/get", utils.HTTP_METHOD_POST, GetUserHttp, true},

	// Team endpoints
	{"team/create", utils.HTTP_METHOD_POST, CreateTeamHttp, true},
	{"team/edit", utils.HTTP_METHOD_POST, EditTeamHttp, true},
	{"team/edit/owner", utils.HTTP_METHOD_POST, EditTeamOwnerHttp, true},
	{"team/delete", utils.HTTP_METHOD_POST, DeleteTeamHttp, true},
	{"team/get", utils.HTTP_METHOD_POST, GetTeamHttp, true},

	// Role endpoints
	{"rol/create", utils.HTTP_METHOD_POST, CreateRoleHttp, true},
	{"rol/edit", utils.HTTP_METHOD_POST, EditRoleHttp, true},
	{"rol/delete", utils.HTTP_METHOD_POST, DeleteRoleHttp, true},
	{"rol/get", utils.HTTP_METHOD_POST, GetRoleHttp, true},

	// Ping
	{"ping", utils.HTTP_METHOD_GET, PingHttp, false},
}

func Start() {

	log.Logger.WithDebug()
	log.ShowLogAppTitle()

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.Use(SecurityMiddleware())

	registerEndpoints(router)

	log.FormattedInfo("API started on https://${0}:${1}${2}", configuration.Params.Ip, configuration.Params.Port, API_COMPLETE)
	state := router.Run(configuration.Params.Ip + ":" + configuration.Params.Port)
	log.Error(state.Error())

}

func registerEndpoints(router *gin.Engine) {

	for _, endpoint := range endpoints {
		switch endpoint.Method {
		case utils.HTTP_METHOD_GET:
			router.GET(API_COMPLETE+endpoint.Path, endpoint.Listener)
		case utils.HTTP_METHOD_POST:
			router.POST(API_COMPLETE+endpoint.Path, endpoint.Listener)
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SecurityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Check if endpoint is secured
		for _, endpoint := range endpoints {
			if API_COMPLETE+endpoint.Path == c.Request.URL.Path {
				if !endpoint.Secured {
					log.FormattedInfo("Endpoint ${0} is not secured", endpoint.Path)
					return
				}
			}
		}

		log.FormattedInfo("Endpoint ${0} is secured", c.Request.URL.Path)

		// Get token
		token := c.Request.Header.Get("Authorization")

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
		err := IsTokenValid(client, token)

		if err != nil {
			c.AbortWithStatusJSON(
				err.Code,
				gin.H{"code": err.Error, "message": err.Message},
			)

			return
		}
	}
}
