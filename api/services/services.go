package services

import (
	"github.com/akrck02/valhalla-core/configuration"
	"github.com/akrck02/valhalla-core/log"
	"github.com/gin-gonic/gin"
)

const API_PATH = "api"
const VERSION = "v1"
const API_COMPLETE = "/" + API_PATH + "/" + VERSION + "/"

func Start() {

	log.Logger.WithDebug()
	log.ShowLogAppTitle()

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET(API_COMPLETE+"ping", PingHttp)

	// User endpoints
	router.POST(API_COMPLETE+"user/register", RegisterHttp)
	router.POST(API_COMPLETE+"user/login", LoginHttp)
	router.POST(API_COMPLETE+"user/edit", EditUserHttp)
	router.POST(API_COMPLETE+"user/edit/email", EditUserEmailHttp)
	router.POST(API_COMPLETE+"user/delete", DeleteUserHttp)
	router.POST(API_COMPLETE+"user/get", GetUserHttp)

	// Team endpoints
	router.POST(API_COMPLETE+"team/create", CreateTeamHttp)
	router.POST(API_COMPLETE+"team/edit", EditTeamHttp)
	router.POST(API_COMPLETE+"team/delete", DeleteTeamHttp)
	router.POST(API_COMPLETE+"team/get", GetTeamHttp)

	log.FormattedInfo("API started on https://${0}:${1}${2}", configuration.Params.Ip, configuration.Params.Port, API_COMPLETE)
	state := router.Run(configuration.Params.Ip + ":" + configuration.Params.Port)
	log.Error(state.Error())

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
