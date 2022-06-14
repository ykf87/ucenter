package user

import (
	"reflect"
	"strings"
	"ucenter/app/coder/mailcode"
	"ucenter/app/config"
	"ucenter/app/controllers"
	"ucenter/app/safety/passwordhash"
	"ucenter/models"

	"github.com/gin-gonic/gin"
)

//获取邮箱验证码
func Emailcode(c *gin.Context) {
	email := c.PostForm("email")

	var lang string
	langos, exits := c.Get("_lang")
	if !exits {
		lang = config.Config.Lang
	} else {
		lang = langos.(string)
	}
	err := models.GetEmailCode(email, lang)
	if err != nil {
		controllers.Error(c, nil, &controllers.Msg{Str: err.Error()})
		return
	}
	controllers.Success(c, nil, &controllers.Msg{Str: "Captcha sent successfully, please check your email"})
}

//注册
func Sign(c *gin.Context) {
	account := c.PostForm("account")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	pwd := c.PostForm("password")
	invite := c.PostForm("invite")
	code := c.PostForm("code")

	ip := c.ClientIP()
	user, err := models.MakeUser(account, email, phone, pwd, code, invite, ip)
	if err != nil {
		controllers.Error(c, nil, &controllers.Msg{Str: err.Error()})
		return
	}
	if user == nil {
		controllers.Error(c, nil, &controllers.Msg{Str: "System error, please try again later"})
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
	code := c.PostForm("code")
	veried := false

	if email != "" && code != "" {
		err := mailcode.Verify(email, code)
		if err != nil {
			controllers.Error(c, nil, &controllers.Msg{Str: err.Error()})
			return
		}
		veried = true
	}

	user := models.GetUser(0, account, email, phone)
	var msg string
	if user != nil {
		if veried == true {
			token := user.Token()
			if token != "" {
				controllers.Success(c, map[string]interface{}{
					"token": token,
				}, nil)
				return
			} else {
				msg = "Voucher generation failed, please try again later"
			}
		} else if user.Pwd != "" {
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

//编辑信息
func Editer(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)
	key := c.PostForm("k")
	val := c.PostForm("v")
	if key == "" || val == "" {
		controllers.Error(c, nil, &controllers.Msg{Str: "Missing editorial content"})
		return
	}

	inputs := make([]reflect.Value, 2)
	inputs[0] = reflect.ValueOf(user)
	inputs[1] = reflect.ValueOf(val)
	tf := reflect.TypeOf(user.Edinfo)
	vl := reflect.ValueOf(user.Edinfo)
	key = strings.ToUpper(string(key[0])) + key[1:]
	mth, ok := tf.MethodByName("Set" + key)
	if ok == true {
		rs := vl.Method(mth.Index).Call(inputs)
		err := rs[0].Interface()
		if err != nil {
			controllers.Error(c, nil, &controllers.Msg{Str: err.(error).Error()})
		} else {
			var dt map[string]interface{}
			if len(rs) > 1 {
				if rs[1].CanInterface() == true {
					dt = rs[1].Interface().(map[string]interface{})
				}
			}
			controllers.Success(c, dt, &controllers.Msg{Str: "Success"})
		}
	} else {
		controllers.Error(c, nil, &controllers.Msg{Str: "No modification allowed"})
	}
}

//忘记密码
func Forgot(c *gin.Context) {
	email := c.PostForm("email")
	code := c.PostForm("code")
	pwd := c.PostForm("password")
	msg := "System error, please try again later"

	if email == "" {
		msg = "Please set the Email"
	} else if code == "" {
		msg = "Please input your Captcha"
	} else if pwd == "" {
		msg = "Please set a password"
	} else {
		err := mailcode.Verify(email, code)
		if err != nil {
			msg = err.Error()
		} else {
			user := models.GetUser(0, "", email, "")
			if user == nil {
				msg = "Account not found"
			} else {
				err := user.ChangePwd(pwd)
				if err != nil {
					msg = err.Error()
				} else {
					token := user.Token()
					if token != "" {
						controllers.Success(c, map[string]interface{}{"token": token}, &controllers.Msg{Str: "Success"})
						return
					}
				}
			}
		}
	}
	controllers.Error(c, nil, &controllers.Msg{Str: msg})
}
