package user

import (
	"ucenter/app/controllers"
	"ucenter/app/safety/passwordhash"
	"ucenter/models"

	"github.com/gin-gonic/gin"
)

//注册
func Sign(c *gin.Context) {
	account := c.PostForm("account")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	pwd := c.PostForm("password")
	// code := c.PostForm("code")

	user, err := models.MakeUser(account, email, phone, pwd, c.ClientIP())
	if err != nil {
		controllers.Error(c, nil, &controllers.Msg{Str: err.Error()})
		return
	}
	token := user.Token()
	if token == "" {
		controllers.Error(c, nil, &controllers.Msg{Str: "Voucher generation failed, please try again later"})
		return
	}
	controllers.Success(c, map[string]interface{}{
		"token": token,
	}, nil)
	return
}

//登录
func Login(c *gin.Context) {
	account := c.PostForm("account")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	pwd := c.PostForm("password")

	user := models.GetUser(0, account, email, phone)
	var msg string
	if user != nil {
		if user.Pwd != "" {
			if passwordhash.PasswordVerify(pwd, user.Pwd) != true {
				msg = "Password error"
			} else {
				if user.Status != 1 {
					msg = "Your account is abnormal"
				} else {
					token := user.Token()
					if token != "" {
						controllers.Success(c, map[string]interface{}{
							"token": token,
						}, nil)
						return
					} else {
						msg = "Voucher generation failed, please try again later"
					}
				}
			}
		} else if pwd != "" {
			msg = "Account cannot be logged in with password,password not set"
		} else {
			msg = "Incomplete login information"
		}
	} else {
		msg = "Account not found"
	}
	controllers.Error(c, nil, &controllers.Msg{Str: msg})
}

//个人信息
func Index(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)
	controllers.Success(c, user.Info(), nil)
}
