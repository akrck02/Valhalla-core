package utils

import "github.com/gin-gonic/gin"

func ReadBodyJson(c *gin.Context, obj interface{}) error {
	return c.BindJSON(obj)
}
