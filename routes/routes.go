package routes

import (
	"smart-home-server/api/auth"   // 引用 auth 下的 Register 和 Login 处理函数
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// 用户模块：处理注册和登录
	userGroup := r.Group("/api/user")
	{
		userGroup.POST("/register", auth.Register)  // 调用 auth.Register 函数处理注册
		userGroup.POST("/login", auth.Login)        // 调用 auth.Login 函数处理登录
	}

	// 其他模块（比如：数据查询、设备控制等）可以在这里继续注册
}