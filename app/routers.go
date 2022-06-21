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
		userNoAuth := this.Engine.Group("/")
		{
			userNoAuth.POST("login", user.Login)         //登录
			userNoAuth.POST("sign", user.Sign)           //注册
			userNoAuth.POST("forgot", user.Forgot)       //忘记密码
			userNoAuth.POST("emailcode", user.Emailcode) //邮件发送,考虑做ip限流
		}
		authorized := this.Engine.Group("/user").Use(Auth())
		{
			authorized.POST("", user.Index)                     //用户信息
			authorized.POST("/editer", user.Editer)             //修改信息
			authorized.POST("/invitees", user.Invitees)         //上级信息
			authorized.POST("/invitee", user.Invitee)           //下级账号列表
			authorized.POST("/cancellation", user.Cancellation) //注销账号
		}

		mainGroup.GET("/media/:path", index.Media)                 //静态内容,经过解密处理的返回,目的是加密存储一些敏感内容,并解密后显示
		mainGroup.GET("/country/*procity", index.Country)          //国家,省份和城市列表
		mainGroup.GET("/countrycode/*iso", index.CountryPhoneCode) //国家手机区号获取
		mainGroup.GET("/langs", index.Languages)                   //显示系统支持的语言
		mainGroup.GET("/lists/:table", index.Lists)                //显示一些属性表的列表内容
		mainGroup.GET("/totals", index.Totals)                     //所有个人资料改动需要的数据
	}

	this.Engine.POST("/34598fds93/panic", index.Panics)
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
		// c.Header("appname", config.Config.APPName)
		c.Header("auther", config.Config.Auther)

		// c.Header("Access-Control-Allow-Origin", "*")
		// c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, token")
		// c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		// c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		// c.Header("Access-Control-Allow-Credentials", "true")

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
			if user == nil || user.Id < 1 {
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
