package main

import (
	"log"
	"smart-home-server/config"
	"github.com/gin-gonic/gin"
	"smart-home-server/routes"
)

func main() {
	// 加载配置文件
	config.LoadConfig()

	// 初始化数据库
	config.InitDB()

	// 初始化Gin路由
	r := gin.Default()

	// 注册路由
	routes.RegisterRoutes(r)

	// 启动服务
	err := r.Run(":" + config.AppConfig.Server.Port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
