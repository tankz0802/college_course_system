package controller

import (
	"ccs/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCLassListHandler(c *gin.Context) {
	classList, err := model.GetClassList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": classList,
	})
}
