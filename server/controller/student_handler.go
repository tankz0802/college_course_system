package controller

import (
	"ccs/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStudentListHandler(c *gin.Context) {
	studentList, err := model.GetStudentList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": studentList,
	})
}
