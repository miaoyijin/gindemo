package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RegisterRoute(r *gin.Engine)  {
	r.GET("/ping", func(c *gin.Context) {
		//whileTrue()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	/**
	r.LoadHTMLGlob("./public/html/*")
	 */
	r.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}

func whileTrue()  {
	go func() {
		for{
			fmt.Print(111)
		}
		time.Sleep(1* time.Second)
	}()
}
