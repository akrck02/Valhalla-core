package middleware

import (
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

// Manage not found endpoint requests
//
// [return] func(c *gin.Context): handler
func NotFound() func(c *gin.Context) {

	return func(c *gin.Context) {
		c.JSON(utils.HTTP_STATUS_NOT_FOUND, gin.H{"code": utils.HTTP_STATUS_NOT_FOUND, "message": "Not found"})
	}

}
