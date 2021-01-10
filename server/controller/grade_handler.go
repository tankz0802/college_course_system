package controller

import (
	"ccs/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateGradeHandler(c *gin.Context) {
	grade := &model.Grade{}
	if err := c.ShouldBindJSON(grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	err := grade.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	courseStudentGradeList, _ := model.GetCourseStudentGradeList(grade.Cid)
	c.JSON(http.StatusOK, gin.H{
		"data": courseStudentGradeList,
	})
}

func GetGradeListHandler(c *gin.Context) {
	sid, _ := strconv.ParseInt(c.Param("sid"), 10, 64)
	gradeList, err := model.GetGradeInfoList(sid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gradeList,
	})
}

func GetCourseStudentGradeListHandler(c *gin.Context) {
	cid := c.Param("cid")
	courseStudentGradeList, err := model.GetCourseStudentGradeList(cid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": courseStudentGradeList,
	})
}