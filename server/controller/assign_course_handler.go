package controller

import (
	"ccs/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AssignCourseHandler(c *gin.Context) {
	assignCourse := &model.AssignCourse{}
	c.ShouldBindJSON(assignCourse)
	_, err := assignCourse.Add()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}