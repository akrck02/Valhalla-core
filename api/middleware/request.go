package middleware

import (
	"github.com/akrck02/valhalla-core/models"
	"github.com/gin-gonic/gin"
)

func Request() gin.HandlerFunc {
	return func(c *gin.Context) {

		var request = models.Request{
			Authorization: c.Request.Header.Get(AUTHORITATION_HEADER),
			IP:            c.ClientIP(),
			UserAgent:     c.Request.UserAgent(),
		}

		c.Set("request", request)
		c.Next()
	}
}
