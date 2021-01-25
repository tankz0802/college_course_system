package main

import (
	"ccs/config"
	"ccs/controller"
	"ccs/db"
	"ccs/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func main() {
	gin.DefaultWriter = colorable.NewColorableStdout()
	config.Init()
	fmt.Println("password:", config.MYSQL_PASSWORD)
	db.Init()
	utils.InitData()
	router := gin.Default()
	v1 := router.Group("/ccs")
	{
		v1.POST("/login", controller.LoginHandler)
		v1.GET("/elective_course/:sid", controller.GetElectiveCourseHandler)
		v1.POST("/select_elective_course", controller.SelectElectiveCourseHandler)
		v1.POST("/cancel_elective_course", controller.CancelElectiveCourseHandler)
		v1.GET("/teacher_course_table/:tid", controller.TeacherCourseTableHandler)
		v1.GET("/student_course_table/:sid", controller.StudentCourseTableHandler)
		v1.GET("/assign_course_list/:tid", controller.GetAssignCourseListHandler)
		v1.GET("/teach_course_list/:tid", controller.GetTeachCourseListHandler)
		v1.GET("/class_list", controller.GetCLassListHandler)
		v1.POST("/assign_course", controller.AssignCourseHandler)
		v1.POST("/update_course/:tid", controller.UpdateCourseHandler)
		v1.GET("/grade_list/:sid", controller.GetGradeListHandler)
		v1.GET("/course_student_grade_list/:cid", controller.GetCourseStudentGradeListHandler)
		v1.POST("/update_grade", controller.UpdateGradeHandler)
		v1.GET("/student_list", controller.GetStudentListHandler)
		v1.GET("/teacher_list", controller.GetTeacherListHandler)
		v1.GET("/course_list", controller.GetCourseListHandler)
		v1.GET("/class_info_list", controller.GetClassInfoListHandler)
	}
	fmt.Printf(`
      ___           ___           ___     
     /\__\         /\__\         /\__\    
    /:/  /        /:/  /        /:/ _/_   
   /:/  /        /:/  /        /:/ /\  \  
  /:/  /  ___   /:/  /  ___   /:/ /::\  \ 
 /:/__/  /\__\ /:/__/  /\__\ /:/_/:/\:\__\
 \:\  \ /:/  / \:\  \ /:/  / \:\/:/ /:/  /
  \:\  /:/  /   \:\  /:/  /   \::/ /:/  / 
   \:\/:/  /     \:\/:/  /     \/_/:/  /  
    \::/  /       \::/  /        /:/  /   
     \/__/         \/__/         \/__/    
	`)
	router.Run(":12345")
}