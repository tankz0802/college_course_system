package main

import (
	"ccs/config"
	"ccs/controller"
	"ccs/db"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func main() {
	gin.DefaultWriter = colorable.NewColorableStdout()
	config.Init()
	db.Init()
	//utils.InitData()
	router := gin.Default()
	v1 := router.Group("/ccs")
	{
		v1.POST("/login", controller.LoginHandler)
		v1.GET("/elective_course/:sid", controller.GetElectiveCourseHandler)
		v1.POST("/select_elective_course", controller.SelectElectiveCourseHandler)
		v1.GET("/teacher_course_table/:tid", controller.TeacherCourseTableHandler)
		v1.GET("/student_course_table/:sid", controller.StudentCourseTableHandler)
		v1.GET("/teach_course_list/:tid", controller.GetTeachCourseListHandler)
		v1.GET("/class_list", controller.GetCLassListHandler)
		v1.POST("/assign_course", controller.AssignCourseHandler)
	}
	router.Run(":12345")
}