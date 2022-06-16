package app

import (
	"ucenter/app/config"
	"ucenter/app/controllers"
	"ucenter/app/controllers/index"
	"ucenter/app/controllers/user"
	"ucenter/models"

	"github.com/gin-gonic/gin"
)

func Init() {

}

//web的路由
func (this *AppClient) WebRouter() {
	mainGroup := this.Engine.Use(Middle())
	{
		mainGroup.GET("/media/:path", index.Media)
		mainGroup.GET("/country/*procity", index.Country)
		mainGroup.GET("/lists/:table", index.Lists)
		mainGroup.POST("/login", user.Login)
		mainGroup.POST("/sign", user.Sign)
		mainGroup.POST("/forgot", user.Forgot)
		mainGroup.POST("/emailcode", user.Emailcode)

		authorized := mainGroup.Use(Auth())
		{
			authorized.POST("/index", user.Index)
			authorized.POST("/editer", user.Editer)
		}
	}

}

func Middle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var lang string
		if c.GetHeader("lang") != "" {
			lang = c.GetHeader("lang")
		} else if cc, err := c.Cookie("lang"); err == nil {
			lang = cc
		} else {
			lang = config.Config.Lang
		}
		c.Set("_lang", lang)

		country, err := models.GetCountryByIp(c.ClientIP())
		if err != nil {
			c.Set("_timezone", config.Config.Timezone)
			c.Header("timezone", config.Config.Timezone)
		} else {
			c.Set("_timezone", country.Timezone)
			c.Header("timezone", country.Timezone)
		}

		c.Header("language", lang)
		c.Header("server", config.Config.APPName)
		c.Header("appname", config.Config.APPName)
		c.Header("auther", config.Config.Auther)
		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		token = c.GetHeader("token")
		if token == "" {
			token = c.GetString("token")
		}
		if token == "" {
			controllers.Resp(c, nil, &controllers.Msg{Str: "Please Login"}, 401)
			c.Abort()
		} else {
			user := models.UnToken(token)
			if user == nil {
				controllers.Resp(c, nil, &controllers.Msg{Str: "Please Login"}, 401)
				c.Abort()
			} else {
				c.Set("_user", user)
				if user.Lang != "" {
					c.Header("language", user.Lang)
					c.Set("_lang", user.Lang)
				}
				if user.Timezone != "" {
					c.Set("_timezone", user.Timezone)
				}
				c.Next()
			}
		}
	}
}
