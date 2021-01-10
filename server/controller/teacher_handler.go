package controller

import (
	"ccs/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTeacherListHandler(c *gin.Context) {
	teacherList, err := model.GetTeacherList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": teacherList,
	})
}
