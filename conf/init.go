package conf

import (
	"DuckyGo-MVC/cache"
	"DuckyGo-MVC/middlewares"
	"DuckyGo-MVC/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/go-box/pongo2gin"
	"os"
	"path/filepath"
	"time"
)

func Init() *gin.Engine {
	route := gin.Default()

	// 中间件
	// 没开启redis的时候就用cookie
	if os.Getenv("plugin") == "1" {
		route.Use(middlewares.SessionRedis(os.Getenv("SECRET_KEY")))
	} else {
		route.Use(middlewares.SessionCookie(os.Getenv("SECRET_KEY")))
	}

	// 基础变量 和 加载环境变量
	base, _ := os.Getwd()
	_ = godotenv.Load()

	// 使用pongo2gin替换gin的模板语言 并指定静态文件目录
	route.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{
		TemplateDir: filepath.Join(base, "templates"),
		ContentType: "text/html; charset=utf-8",
	})

	// 静态文件配置
	route.Static("/static", filepath.Join(base, "static"))

	// 网站favicon
	route.StaticFile("/favicon.ico", filepath.Join(base, "static", "favicon.ico"))

	// 加载路由
	MainUrls(route.Group("/"))

	if os.Getenv("plugin") == "1" {
		// 数据库连接 （MySQL
		models.Init()

		// 缓存连接 （Redis
		cache.Init()
	}

	// 人性化提示服务器开启成功！
	go func() {
		time.Sleep(1 * time.Second)
		info := `==================== DuckyGo-MVC 启动成功！！！！！！！ ==================== `
		fmt.Println(info)
	}()

	return route
}
