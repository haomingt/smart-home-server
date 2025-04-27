package main

import (
	"log"
	"smart-home-server/config"
	"github.com/gin-gonic/gin"
	"smart-home-server/routes"
	"github.com/robfig/cron/v3"
	"smart-home-server/api/monitoring"
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
	// 创建一个新的 Cron 调度器
	c := cron.New(cron.WithSeconds()) // 设置 Cron 支持秒级别的调度

	// 定义一个 Cron 表达式，表示每分钟执行一次
	_, err = c.AddFunc("0 * * * * *", func() { // 每分钟执行
		// 获取并更新宠物信息
		if err := monitoring.FetchPetInfoFromOtherGroup(); err != nil {
			log.Printf("Failed to fetch pet info: %v", err)
		} else {
			log.Println("Successfully fetched and updated pet info")
		}

		// 获取并更新喂食计划
		if err := monitoring.FetchFeedingPlanFromOtherGroup(); err != nil {
			log.Printf("Failed to fetch feeding plan: %v", err)
		} else {
			log.Println("Successfully fetched and updated feeding plan")
		}

		// 获取并更新燥剂更换提醒
		if err := monitoring.FetchMedicationReminderFromOtherGroup(); err != nil {
			log.Printf("Failed to fetch medication reminder: %v", err)
		} else {
			log.Println("Successfully fetched and updated medication reminder")
		}

		// 获取并更新剩余粮食量
		if err := monitoring.FetchPetFoodInventoryFromOtherGroup(); err != nil {
			log.Printf("Failed to fetch food inventory: %v", err)
		} else {
			log.Println("Successfully fetched and updated food inventory")
		}
	})

	if err != nil {
		log.Fatalf("Error adding cron task: %v", err)
	}

	// 启动 Cron 调度器
	c.Start()

	// 为了让应用持续运行，保持主线程不退出
	select {}
}
