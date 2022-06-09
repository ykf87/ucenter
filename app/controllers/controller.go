package controllers

import (
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}, msg string, code int) {
	if code == 0 {
		code = 200
	}

	c.JSON(200, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func Error(c *gin.Context, data interface{}, msg string, code int) {
	if code == 0 {
		code = 500
	}

	c.JSON(200, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
