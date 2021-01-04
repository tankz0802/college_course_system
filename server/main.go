package main

import "github.com/gin-gonic/gin"
import "ccs/controller"

func main() {
	router := gin.Default()
	v1 := router.Group("/api/student")
	{
		v1.POST("/login", controller.LoginHandler)
		v1.POST("/register", controller.RegisterHandler)
	}
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}