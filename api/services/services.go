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

	var configuration = configuration.LoadConfiguration()

	log.ShowLogAppTitle()

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET(API_COMPLETE+"ping/", Ping)
	router.POST(API_COMPLETE+"register/", Register)
	router.POST(API_COMPLETE+"login/", Login)

	log.FormattedInfo("API started on https://${0}:${1}${2}", configuration.Ip, configuration.Port, API_COMPLETE)
	state := router.Run(configuration.Ip + ":" + configuration.Port)
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
