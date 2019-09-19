package main

import (
	"DuckyGo-MVC/conf"
)

func main() {
	// 网站本身
	engine := conf.Init()

	// 运行  8001
	engine.Run(":8001")
}
