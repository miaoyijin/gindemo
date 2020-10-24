package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoute(r *gin.Engine)  {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.LoadHTMLGlob("./public/html/*")
	r.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
