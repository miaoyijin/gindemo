package main;

import (
	"github.com/gin-gonic/gin"
	"tools/routers"
)

func main() {
	r := gin.Default()
	routers.RegisterRoute(r);
	r.Run("0.0.0.0:8081")
}