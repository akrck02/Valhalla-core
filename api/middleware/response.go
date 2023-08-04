package middleware

import (
	"time"

	"github.com/akrck02/valhalla-core/models"
	"github.com/gin-gonic/gin"
)

// Manage errors in a generic way passing the function that will be executed
//
// [param] endpoint | Endpoint: endpoint
// [return] func(c *gin.Context): handler
func APIResponseManagement(endpoint models.Endpoint) func(c *gin.Context) {

	return func(c *gin.Context) {

		//calculate the time of the request
		start := time.Now()
		result, error := endpoint.Listener(c)
		end := time.Now()
		elapsed := end.Sub(start)

		if error != nil {
			// log.FormattedError("Error ${0} : ${1} in ${2}", error.Error, error.Message, c.Request.URL.Path)
			c.JSON(error.Status, error)
			return
		}

		result.ResponseTime = elapsed.Nanoseconds()
		c.JSON(result.Code, result)
	}

}
