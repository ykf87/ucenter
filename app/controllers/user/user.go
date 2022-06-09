package user

import (
	"ucenter/app/controllers"

	"github.com/gin-gonic/gin"
)

//注册
func Sigin(c *gin.Context) {

}

//登录
func Login(c *gin.Context) {
	controllers.Error(c, nil, "登录失败", 0)
}

//主页
func Index(c *gin.Context) {
	controllers.Success(c, nil, "", 0)
}
