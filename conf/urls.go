package conf

import (
	"DuckyGo-MVC/controlls"
	"github.com/gin-gonic/gin"
)

func Urls(r *gin.RouterGroup) {
	r.GET("/", controlls.Index)

}
