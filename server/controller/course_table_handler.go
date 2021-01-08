package controller

import (
	"ccs/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TeacherCourseTableHandler(c *gin.Context) {
	tid, _ := strconv.ParseInt(c.Param("tid"), 10, 64)
	courseTable, err := model.GetTeacherCourseTable(tid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": courseTable,
	})
}

func StudentCourseTableHandler(c *gin.Context) {
	sid, _ := strconv.ParseInt(c.Param("sid"), 10, 64)
	courseTable, err := model.GetStudentCourseTable(sid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": courseTable,
	})
}