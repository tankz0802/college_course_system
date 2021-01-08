package controller

import (
	"ccs/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AssignCourseHandler(c *gin.Context) {
	assignCourse := &model.AssignCourse{}
	c.ShouldBindJSON(assignCourse)
	fmt.Println(assignCourse)
	return
}