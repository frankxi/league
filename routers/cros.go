package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cros(r *gin.Engine) {
	//TODO config option cors
	r.Use(cors.Default())
}
