package controllers

import (
	"html/template"
	"ucenter/app/i18n"

	"github.com/gin-gonic/gin"
)

type Msg struct {
	Str  string
	Args []interface{}
}

func Success(c *gin.Context, data interface{}, msg *Msg) {
	Resp(c, data, msg, 200)
}

func Error(c *gin.Context, data interface{}, msg *Msg) {
	Resp(c, data, msg, 500)
}

func Resp(c *gin.Context, data interface{}, msg *Msg, code int) {
	if code == 0 {
		code = 200
	}

	langobj, _ := c.Get("_lang")
	lang := langobj.(string)

	var msgStr template.HTML
	if msg != nil {
		if msg.Args == nil {
			msgStr = i18n.T(lang, msg.Str)
		} else {
			msgStr = i18n.T(lang, msg.Str, msg.Args...)
		}
	}
	c.JSON(200, map[string]interface{}{
		"code": code,
		"msg":  msgStr,
		"data": data,
	})
}
