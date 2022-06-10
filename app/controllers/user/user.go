package user

import (
	"fmt"
	"ucenter/app/controllers"

	"github.com/gin-gonic/gin"
)

//注册
func Sigin(c *gin.Context) {
	controllers.Error(c, nil, &controllers.Msg{Str: "Sign failed"})
}

//登录
func Login(c *gin.Context) {
	account := c.PostForm("account")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	pwd := c.PostForm("password")

	fmt.Println(account, phone, email, pwd, "--------------")
	controllers.Error(c, nil, &controllers.Msg{Str: "Login failed" + account + "---"})
}

//主页
func Index(c *gin.Context) {
	controllers.Success(c, nil, nil)
}
