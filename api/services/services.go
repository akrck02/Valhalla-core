package services

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/withmandala/go-log"
	"go.mongodb.org/mongo-driver/mongo"
)

const API_PATH = "api"
const VERSION = "v1"
const API_COMPLETE = "/" + API_PATH + "/" + VERSION + "/"

func Start(logger *log.Logger, conn context.Context, client mongo.Client) {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET(API_COMPLETE+"ping/", Ping)
	router.GET(API_COMPLETE+"register/", route(Register, logger, conn, client))

	logger.Info("Server started on 127.0.0.1:3333")
	state := router.Run("127.0.0.1:3333")

	logger.Error(state)

}

func route(function func(c *gin.Context, logger *log.Logger, conn context.Context, client mongo.Client), logger *log.Logger, conn context.Context, client mongo.Client) func(c *gin.Context) {

	return func(c *gin.Context) {
		function(c, logger, conn, client)
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
