package controller

import (
	"ccs/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LoginData struct {
	ID string `json:"id" form:"id"`
	Password string `json:"password" form:"password"`
}

func LoginHandler(c *gin.Context) {
	var loginData LoginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	var role,name string
	id , _ := strconv.ParseInt(loginData.ID, 10 ,64)
	if len(loginData.ID) == 6 {
		teacher := &model.Teacher{
			Id: id,
			Password: loginData.Password,
		}
		if !teacher.TeacherIsExists() {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "用户不存在",
			})
			return
		}
		if !teacher.CheckPassword() {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "密码错误",
			})
			return
		}
		name = teacher.Name
		if teacher.Title == 1 {
			role = "system"
		}else{
			role = "teacher"
		}
	}else if len(loginData.ID) == 10 {
		student := &model.Student{
			Id: id,
			Password: loginData.Password,
		}
		if !student.UserIsExists() {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "用户不存在",
			})
			return
		}
		if !student.CheckPassword() {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "密码错误",
			})
			return
		}
		name = student.Name
		role = "student"
	}else{
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"role": role,
		"name": name,
	})
}