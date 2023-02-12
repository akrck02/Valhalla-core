package utils

import "github.com/gin-gonic/gin"

type HttpStatusCode int

const (
	HTTP_STATUS_OK                         = 200
	HTTP_STATUS_CREATED                    = 201
	HTTP_STATUS_ACCEPTED                   = 202
	HTTP_STATUS_NO_CONTENT                 = 204
	HTTP_STATUS_BAD_REQUEST                = 400
	HTTP_STATUS_UNAUTHORIZED               = 401
	HTTP_STATUS_FORBIDDEN                  = 403
	HTTP_STATUS_NOT_FOUND                  = 404
	HTTP_STATUS_METHOD_NOT_ALLOWED         = 405
	HTTP_STATUS_NOT_ACCEPTABLE             = 406
	HTTP_STATUS_CONFLICT                   = 409
	HTTP_STATUS_INTERNAL_SERVER_ERROR      = 500
	HTTP_STATUS_NOT_IMPLEMENTED            = 501
	HTTP_STATUS_BAD_GATEWAY                = 502
	HTTP_STATUS_SERVICE_UNAVAILABLE        = 503
	HTTP_STATUS_GATEWAY_TIMEOUT            = 504
	HTTP_STATUS_HTTP_VERSION_NOT_SUPPORTED = 505
)

func SendResponse(c *gin.Context, status int, response gin.H) {
	c.JSON(status, response)
}
