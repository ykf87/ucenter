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
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	langos, _ := c.Get("_lang")
	lang := langos.(string)
	timezones, _ := c.Get("_timezone")
	timezone := timezones.(string)

	list := models.GetUserLikedList(user.Id, nil, page, limit)
	for _, v := range list {
		u := models.GetUser(v.Likeid, "", "", "")
		if u != nil && u.Id > 0 {
			v.Info = u.Info(lang, timezone)
		}
	}

	controllers.Success(c, list, &controllers.Msg{Str: "Success"})
}

//喜欢当前用户的列表
func Liker(c *gin.Context) {

}

//相互喜欢列表
func Likes(c *gin.Context) {

}

//取消喜欢
func Unlike(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)
	// langos, _ := c.Get("_lang")
	// lang := langos.(string)
	// timezones, _ := c.Get("_timezone")
	// timezone := timezones.(string)

	id, _ := strconv.Atoi(c.PostForm("id"))

	err := models.UnlikeUser(user.Id, id)
	if err != nil {
		controllers.ErrorNoData(c, err.Error())
		return
	}
	controllers.SuccessStr(c, nil, "Success")
}
