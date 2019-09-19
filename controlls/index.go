package controlls

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index 主页
func Index(c *gin.Context)  {
	/*
		写入逻辑代码 就不分services了.  需要的话自己创建
	*/

	c.HTML(http.StatusOK, "index.html", pongo2.Context{
		"msg": "Index Page",
	})
}