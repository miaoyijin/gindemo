package cmd;

import (
	"flag"
	version "gindemo/common/helpers"
	"gindemo/routers"
	"github.com/gin-qonic/gin"
	"gindemo/helpers"
)
var v *bool
func init()  {
	v = flag.Bool("v", false, "version")
}
func main() {
	flag.Parse()
	if *v {
		version.Version()
		return
	}
	r := gin.Default()
	routers.RegisterRoute(r);
	r.Run("0.0.0.0:8081")
}