package app

import (
	"strings"
	"ucenter/app/config"
	"ucenter/app/controllers"
	"ucenter/app/controllers/albums"
	"ucenter/app/controllers/article"
	"ucenter/app/controllers/index"
	"ucenter/app/controllers/user"
	"ucenter/app/controllers/userlikes"
	"ucenter/models"

	"github.com/gin-gonic/gin"
)

func Init() {

}

//web的路由
func (this *AppClient) WebRouter() {
	mainGroup := this.Engine.Use(Middle())
	{
		articleRouter := this.Engine.Group("/article")
		{
			articleRouter.GET("/:key", article.Index) //文章详情
		}

		//和用户相关的不需要验证权限的接口
		userNoAuth := this.Engine.Group("/")
		{
			userNoAuth.POST("login", user.Login)         //登录
			userNoAuth.POST("sign", user.Sign)           //注册
			userNoAuth.POST("forgot", user.Forgot)       //忘记密码
			userNoAuth.POST("emailcode", user.Emailcode) //邮件发送,考虑做ip限流
		}

		//相册相关
		albumsRoute := this.Engine.Group("/user/albums").Use(Auth())
		{
			albumsRoute.GET("", albums.Albums)                  //公共相册列表
			albumsRoute.GET("/private/*id", albums.Private)     //私密相册列表
			albumsRoute.GET("/list/:id", albums.Albums)         //他人相册列表
			albumsRoute.POST("", albums.UploadAlb)              //上传相册
			albumsRoute.POST("/base64", albums.UploadAlbBase64) //上传相册
			albumsRoute.POST("/remove", albums.Remove)          //删除照片
			albumsRoute.POST("/exg", albums.AlbumsExg)          //相册公共私密互转
		}

		authorized := this.Engine.Group("/user").Use(Auth())
		{
			authorized.GET("", user.Index)          //用户信息
			authorized.GET("/info/:id", user.Index) //用户信息
			// authorized.POST("/editer", user.Editer)             //修改信息
			authorized.POST("/editerbatch", user.EditBatch)     //个人信息批量修改
			authorized.POST("/invitees", user.Invitees)         //上级信息
			authorized.POST("/invitee", user.Invitee)           //下级账号列表
			authorized.POST("/cancellation", user.Cancellation) //注销账号

			authorized.POST("/imsigna", user.Signa) //Im签名

			//喜欢相关
			authorized.POST("/like", userlikes.Like)   //喜欢一个人
			authorized.POST("/liked", userlikes.Liked) //用户喜欢列表
			authorized.POST("/liker", userlikes.Liker) //喜欢当前用户的列表
			authorized.POST("/likes", userlikes.Likes) //相互喜欢列表
		}

		mainGroup.GET("/media/:path", index.Media)                 //静态内容,经过解密处理的返回,目的是加密存储一些敏感内容,并解密后显示
		mainGroup.GET("/country/*iso", index.Country)              //国家,省份和城市列表
		mainGroup.GET("/countrycode/*iso", index.CountryPhoneCode) //国家手机区号获取
		mainGroup.GET("/langs", index.Languages)                   //显示系统支持的语言
		mainGroup.GET("/lists/:table", index.Lists)                //显示一些属性表的列表内容
		mainGroup.GET("/totals", index.Totals)                     //所有个人资料改动需要的数据

		mainGroup.GET("/search", index.Search) //搜索用户
	}

	this.Engine.POST("/34598fds93/panic", index.Panics)
}

func Middle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// signature := c.GetHeader("signature")
		// if signature == "" {
		// 	c.AbortWithStatus(http.StatusNotFound)
		// 	return
		// }
		var lang string
		if cc, err := c.GetQuery("lang"); err == true {
			lang = cc
		} else if c.GetHeader("lang") != "" {
			lang = c.GetHeader("lang")
		} else if cc, err := c.Cookie("lang"); err == nil {
			lang = cc
		} else if c.GetHeader("Accept-Language") != "" {
			langs := strings.Split(c.GetHeader("Accept-Language"), ",")
			lang = langs[0]
		} else {
			lang = config.Config.Lang
		}
		lang = strings.ToLower(lang)
		if _, ok := models.LangLists[lang]; !ok {
			lang = strings.ToLower(config.Config.Lang)
		}
		c.Set("_lang", lang)

		country, err := models.GetCountryByIp(c.ClientIP())
		if err != nil {
			c.Set("_timezone", config.Config.Timezone)
			// c.Header("timezone", config.Config.Timezone)
		} else {
			c.Set("_timezone", country.Timezone)
			// c.Header("timezone", country.Timezone)
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
		user := models.GetUserFromRequest(c)
		if user == nil || user.Id < 1 {
			controllers.Resp(c, nil, &controllers.Msg{Str: "Please Login"}, 401)
			c.Abort()
		} else {
			c.Set("_user", user)
			if user.Timezone != "" {
				c.Set("_timezone", user.Timezone)
			}
			c.Next()
		}
	}
}
