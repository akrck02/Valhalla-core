package middleware

import (
	"github.com/akrck02/valhalla-core/error"
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

// Manage errors in a generic way
//
// [return] gin.HandlerFunc: handler
func Panic() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {

			var err = c.Errors[0]

			c.JSON(utils.HTTP_STATUS_INTERNAL_SERVER_ERROR, &models.Error{
				Code:    utils.HTTP_STATUS_INTERNAL_SERVER_ERROR,
				Error:   error.UNEXPECTED_ERROR,
				Message: err.Error(),
			})
		}
	}
}
