package conf

import (
	"DuckyGo-MVC/controlls"
	"github.com/gin-gonic/gin"
)

func MainUrls(r *gin.RouterGroup) {
	// 主页
	r.GET("/", controlls.Index)

}
