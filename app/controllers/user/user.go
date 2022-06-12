package user

import (
	"fmt"
	"reflect"
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

//编辑信息
func Editer(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)
	key := c.PostForm("k")
	val := c.PostForm("v")

	inputs := make([]reflect.Value, 1)
	inputs[0] = reflect.ValueOf(val)
	tf := reflect.TypeOf(user.Edinfo)
	vl := reflect.ValueOf(user.Edinfo)
	mth, ok := tf.MethodByName(key)
	if ok == true {
		vl.Method(mth.Index).Call(inputs)
		// tf.Method(mth.Index)
	} else {
		controllers.Error(c, nil, &controllers.Msg{Str: "No modification allowed"})
	}
	fmt.Println("================")
	// fmt.Println(mth, ok, "===========")
	// fmt.Println(vl.Method(0).Name)

	// mthlen := vl.Elem().NumMethod()
	// fmt.Println(mthlen, "======")
	// for i := 0; i < mthlen; i++ {
	// 	fmt.Println(vl.Elem().Method(i))
	// }
}
