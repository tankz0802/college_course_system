package controller

import (
	"ccs/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateCourseHandler(c *gin.Context) {
	course := &model.Course{}
	tid, _ := strconv.ParseInt(c.Param("tid"), 10, 64)
	if err := c.ShouldBindJSON(course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	err := course.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	courseList, _ := model.GetUnTeachCourseList(tid)
	c.JSON(http.StatusOK, gin.H{
		"data": courseList,
	})
}

func GetAssignCourseListHandler(c *gin.Context) {
	tid, _ := strconv.ParseInt(c.Param("tid"), 10, 64)
	unTeachCourseList, err := model.GetUnTeachCourseList(tid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	teachCourseList, err := model.GetTeachCourseList(tid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"teach_course_list": teachCourseList,
		"un_teach_course_list": unTeachCourseList,
	})
}

func GetTeachCourseListHandler(c *gin.Context) {
	tid, _ := strconv.ParseInt(c.Param("tid"), 10, 64)
	teachCourseList, err := model.GetTeachCourseList(tid)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": teachCourseList,
	})
}

func GetCourseListHandler(c *gin.Context) {
	courseList, err := model.GetCourseList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": courseList,
	})
}