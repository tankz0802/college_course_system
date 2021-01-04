package main

import (
	"ccs/config"
	"ccs/controller"
	"ccs/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	config.Init()
	router := gin.Default()
	v1 := router.Group("/api/student")
	{
		v1.POST("/login", controller.StudentLoginHandler)
		v1.POST("/register", controller.StudentRegisterHandler)
	}
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}