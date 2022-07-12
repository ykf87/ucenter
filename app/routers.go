package app

import (
	"fmt"
	"ucenter/app/config"
	"ucenter/app/controllers"
	"ucenter/app/controllers/albums"
	"ucenter/app/controllers/article"
	"ucenter/app/controllers/index"
	"ucenter/app/controllers/user"
	"ucenter/app/controllers/userlikes"
	"ucenter/app/rdsmps"
	"ucenter/models"

	"github.com/gin-gonic/gin"
)

func Init() {

}

//web的路由
func (this *AppClient) WebRouter() {
	mainGroup := this.Engine.Group("")
	mainGroup.Use(Api())
	{
		//和用户相关的不需要验证权限的接口
		userNoAuth := mainGroup.Group("/")
		{
			llrr := userNoAuth.Group("")
			{
				llrr.Use(LoginErr()).POST("login", user.Login) //登录
			}
			userNoAuth.POST("sign", user.Sign)                            //注册
			userNoAuth.POST("forgot", user.Forgot)                        //忘记密码
			userNoAuth.Use(Iplimiter()).POST("emailcode", user.Emailcode) //邮件发送,考虑做ip限流
		}

		//需要登录的接口
		mustLoginRouter := mainGroup.Group("/user")
		mustLoginRouter.Use(Auth())
		{
			albumsAuters := mustLoginRouter.Group("/albums")
			{
				albumsAuters.GET("", albums.Albums)                  //公共相册列表
				albumsAuters.GET("/private/*id", albums.Private)     //私密相册列表
				albumsAuters.GET("/list/:id", albums.Albums)         //他人相册列表
				albumsAuters.POST("", albums.UploadAlb)              //上传相册
				albumsAuters.POST("/base64", albums.UploadAlbBase64) //上传相册
				albumsAuters.POST("/remove", albums.Remove)          //删除照片
				albumsAuters.POST("/exg", albums.AlbumsExg)          //相册公共私密互转
			}

			mustLoginRouter.GET("", user.Index)             //用户信息
			mustLoginRouter.GET("/info/:id", user.Index)    //用户信息
			mustLoginRouter.GET("/invitee", user.Invitee)   //上级信息
			mustLoginRouter.GET("/invitees", user.Invitees) //下级账号列表
			// authorized.POST("/editer", user.Editer)             //修改信息-弃用
			mustLoginRouter.POST("/editerbatch", user.EditBatch)     //个人信息批量修改
			mustLoginRouter.POST("/cancellation", user.Cancellation) //注销账号

			mustLoginRouter.POST("/imsigna", user.Signa) //Im签名

			//喜欢相关
			mustLoginRouter.POST("/like", userlikes.Like)   //喜欢一个人
			mustLoginRouter.POST("/liked", userlikes.Liked) //用户喜欢列表
			mustLoginRouter.POST("/liker", userlikes.Liker) //喜欢当前用户的列表
			mustLoginRouter.POST("/likes", userlikes.Likes) //相互喜欢列表
		}

		mainGroup.GET("/media/:path", index.Media)                 //静态内容,经过解密处理的返回,目的是加密存储一些敏感内容,并解密后显示
		mainGroup.GET("/country/*cid", index.Country)              //国家,省份和城市列表
		mainGroup.GET("/countrycode/*iso", index.CountryPhoneCode) //国家手机区号获取
		mainGroup.GET("/langs", index.Languages)                   //显示系统支持的语言
		mainGroup.GET("/lists/:table", index.Lists)                //显示一些属性表的列表内容
		mainGroup.GET("/totals", index.Totals)                     //所有个人资料改动需要的数据

		mainGroup.GET("/search", index.Search)     //搜索用户
		mainGroup.GET("/positive", index.Positive) //获取活跃用户列表
		mainGroup.GET("/applangs", index.Applangs) //客户端多语言信息
	}

	webRouters := this.Engine.Group("")
	webRouters.Use(Web())
	{
		//首页
		webRouters.GET("", func(c *gin.Context) {
			c.AbortWithStatus(600)
		})

		webRouters.GET("/invitation", index.Invitation) //通过邀请码注册

		articleRouter := webRouters.Group("/article")
		{
			articleRouter.GET("/info/:key", article.Info) //文章详情
		}
	}

	this.Engine.POST("/34598fds93/panic", index.Panics)
}

func Api() gin.HandlerFunc {
	return func(c *gin.Context) {
		if Signature(c) == false {
			c.AbortWithStatus(404)
			return
		}
		Headers(c)
		c.Next()
	}
}

func Web() gin.HandlerFunc {
	return func(c *gin.Context) {
		Headers(c)
		c.Next()
	}
}

func Signature(c *gin.Context) bool {
	// signature := c.GetHeader("signature")
	// if signature == "" {
	// 	c.AbortWithStatus(http.StatusNotFound)
	// 	return
	// }
	return true
}

func Headers(c *gin.Context) {
	lang := models.GetClientLang(c)
	c.Set("_lang", lang)
	c.Header("language", lang)

	country, err := models.GetCountryByIp(c.ClientIP())
	if err != nil {
		c.Set("_timezone", config.Config.Timezone)
	} else {
		c.Set("_timezone", country.Timezone)
	}

	c.Header("server", config.Config.APPName)
	c.Header("auther", config.Config.Auther)

	// method := c.Request.Method
	// c.Header("Access-Control-Allow-Origin", "*")
	// c.Header("Access-Control-Allow-Headers", "*")
	// c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	// c.Header("Access-Control-Expose-Headers", "*")
	// c.Header("Access-Control-Allow-Credentials", "true")

	// //放行所有OPTIONS方法
	// if method == "OPTIONS" {
	// 	c.AbortWithStatus(http.StatusNoContent)
	// }
}

func Middle() gin.HandlerFunc {
	return func(c *gin.Context) {

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
		user := models.GetUserFromRequest(c)
		if user == nil || user.Id < 1 {
			controllers.Resp(c, nil, &controllers.Msg{Str: "Please Login"}, 401)
			c.Abort()
		} else {
			c.Set("_user", user)
			if user.Timezone != "" {
				c.Set("_timezone", user.Timezone)
			}

			go user.CheckUserUseEnvironment(c)

			c.Next()
		}
	}
}

//登录错误限制,防止爆破
func LoginErr() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		if email != "" {
			cli := &rdsmps.Mmpp{Prev: "login/", Timeout: 300}
			_, errtimes, ttl, _, err := cli.Get(email)
			if err == nil {
				if errtimes > 2 {
					if errtimes == 3 {
						ttl = cli.Timeout
						go cli.Set(email, "", 4, true)
					}
					controllers.Resp(c, ttl, &controllers.Msg{Str: "Too many errors, please try again later"}, 500)
					// c.AbortWithStatusJSON(200, gin.H{
					// 	"code": 500,
					// 	"data": ttl,
					// 	"msg":  nil,
					// })
					c.Abort()
					return
				}
			}
		}
		c.Next()
	}
}

//ip限流
func Iplimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
