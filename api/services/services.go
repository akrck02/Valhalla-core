package services

import (
	"github.com/akrck02/valhalla-core/log"
	"github.com/gin-gonic/gin"
)

const API_PATH = "api"
const VERSION = "v1"
const API_COMPLETE = "/" + API_PATH + "/" + VERSION + "/"

func Start() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET(API_COMPLETE+"ping/", Ping)
	router.POST(API_COMPLETE+"register/", route(Register))
	router.POST(API_COMPLETE+"login/", route(Login))

	log.Info("Server started on 127.0.0.1:3333")
	state := router.Run("127.0.0.1:3333")
	log.Error(state.Error())

}

func route(function func(c *gin.Context)) func(c *gin.Context) {

	return func(c *gin.Context) {
		function(c)
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
