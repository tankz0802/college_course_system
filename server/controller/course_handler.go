package controller

import (
	"ccs/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTeachCourseListHandler(c *gin.Context) {
	tid, _ := strconv.ParseInt(c.Param("tid"), 10, 64)
	teachCourseList, err := model.GetTeachCourseList(tid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": teachCourseList,
	})
}
