package main

import (
	"DuckyGo-MVC/conf"
	"fmt"
)

func main() {
	fmt.Println("Check conf files...")

	// 网站本身
	engine := conf.Init()

	// 运行  8001
	engine.Run(":8001")
}
