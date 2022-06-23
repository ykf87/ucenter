package userlikes

import (
	"strconv"
	"ucenter/app/controllers"
	"ucenter/app/i18n"
	"ucenter/app/push/jiguang"
	"ucenter/models"

	"github.com/gin-gonic/gin"
)

//喜欢一个人
func Like(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)
	langos, _ := c.Get("_lang")

	idstr := c.PostForm("id")
	id, _ := strconv.Atoi(idstr)
	if id < 1 {
		controllers.Error(c, nil, &controllers.Msg{Str: "User does not exist"})
		return
	}
	tuser := models.GetUser(int64(id), "", "", "")
	if tuser == nil {
		controllers.Error(c, nil, &controllers.Msg{Str: "User does not exist"})
		return
	}
	if models.UserLikeAdd(user.Id, tuser.Id) == true {
		controllers.Success(c, nil, &controllers.Msg{Str: "Success"})
		go jiguang.Send(tuser.Id, string(i18n.T(langos.(string), "Someone has a crush on you")))
	} else {
		controllers.Error(c, nil, &controllers.Msg{Str: "System error, please try again later"})
	}
}

//用户喜欢列表
func Liked(c *gin.Context) {

}

//喜欢当前用户的列表
func Liker(c *gin.Context) {

}

//相互喜欢列表
func Likes(c *gin.Context) {

}
