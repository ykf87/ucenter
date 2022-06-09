package app

import (
	"fmt"
	"ucenter/app/controllers/user"

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
	this.Engine.POST("/user/sigin", user.Sigin)
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("prev")
		c.Next()
		fmt.Println("next")
	}
}
