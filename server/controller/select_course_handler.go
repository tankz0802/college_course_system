package controller

import (
	"ccs/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetElectiveCourseHandler(c *gin.Context) {
	fmt.Println(c.Param("sid"))
	sid , _ := strconv.ParseInt(c.Param("sid"), 10, 64)
	fmt.Println(sid)
	courseList, err := model.GetElectiveCourseInfo(sid)
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

func SelectElectiveCourseHandler(c *gin.Context) {
	var sc model.StudentCourse
	if err := c.ShouldBindJSON(&sc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	hasConflict, err := model.CourseHasConflict(sc.SId, sc.CId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if hasConflict {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "课程冲突",
		})
		return
	}
	err = sc.Add()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	courseList, _ := model.GetElectiveCourseInfo(sc.SId)
	c.JSON(http.StatusOK, gin.H{
		"data": courseList,
	})
}