package conf

import (
	"DuckyGo-MVC/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/go-box/pongo2gin"
	"os"
	"path/filepath"
)

//var route *gin.Engine

func Init() *gin.Engine {
	route := gin.Default()

	// 基础变量 和 加载环境变量
	base, _ := os.Getwd()
	_ = godotenv.Load()

	route.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{
		TemplateDir: filepath.Join(base, "templates"),
		ContentType: "text/html; charset=utf-8",
	})

	// 静态文件配置
	route.Static("/static", filepath.Join(base, "static"))

	// 网站favicon
	route.StaticFile("/favicon.ico", filepath.Join(base, "static", "favicon.ico"))


	// 加载路由
	Urls(route.Group("/"))


	/*
	func ?(r *gin.RouterGroup) {}    自定义接收函数

	*/

	// 数据库连接
	models.Init()


	return route
}
