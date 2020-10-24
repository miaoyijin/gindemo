package routers

import "github.com/gin-gonic/gin"

func RegisterRoute(r *gin.Engine)  {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
