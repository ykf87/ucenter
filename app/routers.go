package app

import (
	"ucenter/app/controllers"
	"ucenter/app/controllers/user"
	"ucenter/models"

	"github.com/gin-gonic/gin"
)

func Init() {

}

//web的路由
func (this *AppClient) WebRouter() {
	authorized := this.Engine.Group("/user").Use(Auth())
	{
		authorized.POST("", user.Index)
	}
	this.Engine.POST("/user/login", user.Login)
	this.Engine.POST("/user/sign", user.Sign)
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		token = c.GetHeader("token")
		if token == "" {
			token = c.GetString("token")
		}
		if token == "" {
			controllers.Error(c, nil, &controllers.Msg{Str: "Please Login"})
			c.Abort()
		} else {
			user := models.UnToken(token)
			if user == nil {
				controllers.Error(c, nil, &controllers.Msg{Str: "Please Login"})
				c.Abort()
			}
			c.Set("_user", user)
		}
		c.Next()
	}
}
