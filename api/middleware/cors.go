package middleware

import "github.com/gin-gonic/gin"

// Manage CORS
//
// [return] gin.HandlerFunc: handler
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", `User-Agent, Accept, Accept-Language, Accept-Encoding, Referer, Content-type, mode, Origin, Connection, Sec-Fetch-Dest, Sec-Fetch-Mode, Sec-Fetch-Site, Pragma, Cache-Control`)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
