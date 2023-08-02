package middleware

import (
	"github.com/akrck02/valhalla-core/models"
	"github.com/akrck02/valhalla-core/utils"
	"github.com/gin-gonic/gin"
)

// Manage errors in a generic way passing the function that will be executed
//
// [param] endpoint | Endpoint: endpoint
// [return] func(c *gin.Context): handler
func APIResponseManagement(endpoint models.Endpoint) func(c *gin.Context) {

	return func(c *gin.Context) {

		result, error := endpoint.Listener(c)
		if error != nil {
			// log.FormattedError("Error ${0} : ${1} in ${2}", error.Error, error.Message, c.Request.URL.Path)
			c.JSON(error.Code, error)
			return
		}

		c.JSON(utils.HTTP_STATUS_OK, result)
	}

}
