package user

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"ucenter/app/config"
	"ucenter/app/controllers"
	"ucenter/app/funcs"
	"ucenter/app/i18n"
	"ucenter/app/logs"
	"ucenter/app/mails/sender/bye"
	"ucenter/app/mails/sender/coder"
	"ucenter/app/mails/sender/sign"
	"ucenter/app/safety/passwordhash"
	"ucenter/app/uploadfile/images"
	"ucenter/models"

	"github.com/gin-gonic/gin"
	carbon "github.com/golang-module/carbon/v2"
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
	err := coder.Send(email, lang)
	if err != nil {
		controllers.Error(c, nil, &controllers.Msg{Str: err.Error()})
		return
	}
	controllers.Success(c, nil, &controllers.Msg{Str: "Captcha sent successfully, please check your email"})
}

//注册
func Sign(c *gin.Context) {
	// account := c.PostForm("account")
	// phone := c.PostForm("phone")
	email := c.PostForm("email")
	pwd := c.PostForm("password")
	invite := c.PostForm("invite")
	nickname := c.PostForm("nickname")
	code := c.PostForm("code")
	platform := c.GetHeader("platform")
	langos, exit := c.Get("_lang")

	timezoneo, _ := c.Get("_timezone")
	timezone := timezoneo.(string)

	var lang string
	if !exit {
		lang = config.Config.Lang
	} else {
		lang = langos.(string)
	}

	if email == "" {
		controllers.Error(c, nil, &controllers.Msg{Str: "Please set the Email"})
		return
	}

	ip := c.ClientIP()
	user, err := models.MakeUser("", email, "", pwd, code, invite, nickname, platform, ip, timezone)
	if err != nil {
		logs.Logger.Error(err)
		controllers.Error(c, nil, &controllers.Msg{Str: err.Error()})
		return
	}
	if user == nil {
		controllers.Error(c, nil, &controllers.Msg{Str: "System error, please try again later"})
		return
	}
	if user.Mail != "" {
		go sign.Send(user.Mail, lang)
	}

	token := user.Token()
	if token == "" {
		controllers.Error(c, nil, &controllers.Msg{Str: "Voucher generation failed, please try again later"})
		return
	}

	controllers.Success(c, map[string]interface{}{
		"token":     token,
		"id":        user.Id,
		"signature": user.ImSignature(),
	}, nil)
	return
}

//登录
func Login(c *gin.Context) {
	account := "" //c.PostForm("account")
	phone := ""   //c.PostForm("phone")
	email := c.PostForm("email")
	pwd := c.PostForm("password")
	code := c.PostForm("code")
	veried := false

	if email == "" {
		controllers.Error(c, nil, &controllers.Msg{Str: "Please set the Email"})
		return
	}

	if code != "" {
		err := coder.Verify(email, code)
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
							"token":     token,
							"id":        user.Id,
							"signature": user.ImSignature(),
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
	kk := "login:" + email
	cc, ok := coder.Maps.Get(kk)
	if ok {
		cc.Errtimes += 1
	} else {
		coder.Maps.Set(kk, &coder.MailCodeStruct{Errtimes: 1, Sendtime: time.Now().Unix(), Code: ""})
	}
	controllers.Error(c, nil, &controllers.Msg{Str: msg})
}

//个人信息
func Index(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)
	langob, _ := c.Get("_lang")
	lang := langob.(string)
	timezones, _ := c.Get("_timezone")
	timezone := timezones.(string)

	id := c.Param("id")
	if id != "" {
		uid, _ := strconv.Atoi(id)
		us := models.GetUser(int64(uid), "", "", "")
		if us == nil || us.Id < 1 {
			controllers.ErrorNotFound(c)
			return
		}
		go us.AddVisits()
		us.Visits = us.Visits + 1
		user = us
	}

	info := user.Info(lang, timezone)
	albums := models.GetAlbumList(user.Id, 0, 0, false)
	for _, v := range albums {
		v.Fmt(timezone, lang)
	}
	info["albums"] = albums
	info["inviurl"] = funcs.InviUrl(user.Invite)
	controllers.SuccessStr(c, info, "Success")
	// controllers.Success(c, info, &controllers.Msg{Str: "Success"})
}

//编辑信息
// func Editer(c *gin.Context) {
// 	rs, _ := c.Get("_user")
// 	user, _ := rs.(*models.UserModel)
// 	key := c.PostForm("k")
// 	val := c.PostForm("v")
// 	if key == "" {
// 		controllers.Error(c, nil, &controllers.Msg{Str: "Missing editorial content"})
// 		return
// 	}

// 	inputs := make([]reflect.Value, 3)
// 	inputs[0] = reflect.ValueOf(user)
// 	inputs[1] = reflect.ValueOf(val)
// 	inputs[2] = reflect.ValueOf(c)
// 	tf := reflect.TypeOf(user.Edinfo)
// 	vl := reflect.ValueOf(user.Edinfo)
// 	key = strings.ToUpper(string(key[0])) + key[1:]
// 	mth, ok := tf.MethodByName("Set" + key)
// 	if ok == true {
// 		rs := vl.Method(mth.Index).Call(inputs)
// 		err := rs[0].Interface()
// 		if err != nil {
// 			controllers.Error(c, nil, &controllers.Msg{Str: err.(error).Error()})
// 		} else {
// 			var dt map[string]interface{}
// 			if len(rs) > 1 {
// 				if rs[1].CanInterface() == true {
// 					dt = rs[1].Interface().(map[string]interface{})
// 				}
// 			}
// 			controllers.Success(c, dt, &controllers.Msg{Str: "Success"})
// 		}
// 	} else {
// 		controllers.Error(c, nil, &controllers.Msg{Str: "No modification allowed"})
// 	}
// }

//批量编辑信息
func EditBatch(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)
	lango, _ := c.Get("_lang")
	lang := lango.(string)
	timezoneo, _ := c.Get("_timezone")
	timezone := timezoneo.(string)
	cgdata := make(map[string]interface{})
	rsdata := make(map[string]interface{})

	password := c.PostForm("password") //此处密码为修改账号关键信息使用的密码,并非修改账号密码

	//账号修改
	account := c.PostForm("account")
	if account != "" {
		if account != user.Account {
			if len(account) > 60 {
				controllers.Error(c, map[string]string{"col": "account"}, &controllers.Msg{Str: "Please fill in less than {{$1}} words", Args: []interface{}{60}})
				return
			}
			u := new(models.UserModel)
			models.DB.Table("users").Where("account = ?", account).First(u)
			if u.Id > 0 {
				controllers.Error(c, map[string]string{"col": "account"}, &controllers.Msg{Str: "Account name already exists, please use another one"})
				return
			}
			cgdata["account"] = account
			rsdata["account"] = account
		}
	}

	//邮箱修改
	mail := c.PostForm("mail")
	if mail != "" && mail != user.Mail { //修改邮箱,需要原有邮箱验证码或者账号密码
		mailcode := c.PostForm("mailcode")
		if mailcode != "" {
			err := coder.Verify(mail, mailcode)
			if err != nil {
				controllers.Error(c, map[string]string{"col": "mailcode"}, &controllers.Msg{Str: err.Error()})
				return
			}
		} else if password != "" {
			if passwordhash.PasswordVerify(password, user.Pwd) == false {
				controllers.Error(c, map[string]string{"col": "password"}, &controllers.Msg{Str: "Password error"})
				return
			}
		} else {
			controllers.Error(c, map[string]string{"col": "mail"}, &controllers.Msg{Str: "No modification allowed"})
			return
		}

		u := new(models.UserModel)
		models.DB.Table("users").Where("mail = ?", mail).First(u)
		if u.Id > 0 {
			controllers.Error(c, map[string]string{"col": "mail"}, &controllers.Msg{Str: "Email already exists, please use another one"})
			return
		}
		cgdata["mail"] = mail
		rsdata["mail"] = mail
	}

	//电话修改
	phone := c.PostForm("phone")
	if phone != "" && phone != user.Phone {
		// phonecode := c.PostForm("phonecode")//暂不支持手机号验证
		u := new(models.UserModel)
		models.DB.Table("users").Where("phone = ?", phone).First(u)
		if u.Id > 0 {
			controllers.Error(c, map[string]string{"col": "phone"}, &controllers.Msg{Str: "Phone already exists, please use another one"})
			return
		}
		if user.Phone != "" {
			if password != "" {
				if passwordhash.PasswordVerify(password, user.Pwd) == false {
					controllers.Error(c, map[string]string{"col": "password"}, &controllers.Msg{Str: "Password error"})
					return
				}
			} else {
				controllers.Error(c, map[string]string{"col": "password"}, &controllers.Msg{Str: "Please fill in your password to confirm that you changed it yourself"})
				return
			}
		}
		cgdata["phone"] = phone
		rsdata["phone"] = phone
	}

	age := c.PostForm("age")
	if age != "" {
		ageid, _ := strconv.Atoi(age)
		if ageid < 1 || ageid > 160 {
			controllers.Error(c, map[string]string{"col": "age"}, &controllers.Msg{Str: "Wrong age"})
			return
		}
		cgdata["age"] = age
		rsdata["age"] = age
		user.Age = ageid
	}

	//生日修改
	birth := c.PostForm("birth")
	if birth != "" {
		var fmt string
		fmts, ok := config.Config.Timefmts[lang]
		if ok {
			fmt = fmts.Datefmt
		} else {
			fmt = config.Config.Datefmt
		}

		userBirth := carbon.SetTimezone(timezone).CreateFromTimestamp(user.Birth).Carbon2Time().Format(fmt)
		if strings.Contains(birth, userBirth) == false {
			birthUni := carbon.SetTimezone(timezone).Parse(birth).Timestamp()
			if birthUni < 10000 {
				controllers.Error(c, map[string]string{"col": "birth"}, &controllers.Msg{Str: "Wrong date format"})
				return
			}
			// if user.Age < 1 {
			user.Age = int((time.Now().Unix() - birthUni) / 31536000)
			if user.Age > 0 {
				cgdata["age"] = user.Age
				rsdata["age"] = user.Age
			}
			// }
			cgdata["birth"] = birthUni
			rsdata["birth"] = carbon.SetTimezone(timezone).CreateFromTimestamp(birthUni).Carbon2Time().Format(fmt)
		}
	}

	//修改国家
	country := c.PostForm("country")
	if country != "" {
		id, _ := strconv.Atoi(country)
		countryid := int64(id)
		name := models.CountryMap.Get(lang, countryid)
		if name == "" {
			controllers.Error(c, map[string]string{"col": "country"}, &controllers.Msg{Str: "Please set the content to be modified"})
			return
		}
		cgdata["country"] = countryid
		rsdata["country"] = name
		user.Country = countryid
	}

	//修改城市
	city := c.PostForm("city")
	if city != "" {
		if user.Country < 1 {
			controllers.Error(c, map[string]string{"col": "city"}, &controllers.Msg{Str: "Please select your country first"})
			return
		}
		cityid32, _ := strconv.Atoi(city)
		cityid := int64(cityid32)
		if user.City != cityid {
			if cityid < 1 {
				controllers.Error(c, map[string]string{"col": "city"}, &controllers.Msg{Str: "Please set the content to be modified"})
				return
			}
			citylist, err := models.GetCityByCountryId(lang, user.Country)
			if err != nil {
				controllers.Error(c, map[string]string{"col": "city"}, &controllers.Msg{Str: err.Error()})
				return
			}
			if len(citylist) < 1 {
				controllers.Error(c, map[string]string{"col": "city"}, &controllers.Msg{Str: "Please select your country first"})
				return
			}
			var findcity *models.CityModel
			for _, v := range citylist {
				if v.Id == cityid {
					findcity = v
					break
				}
			}
			if findcity == nil {
				controllers.Error(c, map[string]string{"col": "city"}, &controllers.Msg{Str: "The city of your choice was not found"})
				return
			}
			cgdata["city"] = findcity.Id
			rsdata["city"] = findcity.Name
		}
	}

	//修改星座
	constellation := c.PostForm("constellation")
	if constellation != "" {
		id32, _ := strconv.Atoi(constellation)
		id := int64(id32)
		if id != user.Constellation {
			name := models.ConstellationMap.Get(lang, id)
			if name == "" {
				controllers.Error(c, map[string]string{"col": "constellation"}, &controllers.Msg{Str: "Please set reasonably"})
				return
			}
			cgdata["constellation"] = id
			rsdata["constellation"] = name
		}
	}

	//修改教育程度
	edu := c.PostForm("edu")
	if edu != "" {
		id32, _ := strconv.Atoi(edu)
		id := int64(id32)
		if id != user.Edu {
			name := models.EducationMap.Get(lang, id)
			if name == "" {
				controllers.Error(c, map[string]string{"col": "edu"}, &controllers.Msg{Str: "Please set reasonably"})
				return
			}
			cgdata["edu"] = id
			rsdata["edu"] = name
		}
	}

	//修改情感状态
	emotion := c.PostForm("emotion")
	if emotion != "" {
		id32, _ := strconv.Atoi(emotion)
		id := int64(id32)
		if id != user.Edu {
			name := models.EmotionMap.Get(lang, id)
			if name == "" {
				controllers.Error(c, map[string]string{"col": "emotion"}, &controllers.Msg{Str: "Please set reasonably"})
				return
			}
			cgdata["emotion"] = id
			rsdata["emotion"] = name
		}
	}

	//修改身高
	height := c.PostForm("height")
	if height != "" {
		id, _ := strconv.Atoi(height)
		if id > 0 && id != user.Height {
			if id < 1 || id >= 300 {
				controllers.Error(c, map[string]string{"col": "height"}, &controllers.Msg{Str: "Please set the height reasonably"})
				return
			}
			cgdata["height"] = id
			rsdata["height"] = id
		}
	}

	//修改体重
	weight := c.PostForm("weight")
	if weight != "" {
		id, _ := strconv.ParseFloat(weight, 64)
		if id > 0.0 && id != user.Weight {
			if id < 1.0 || id >= 500000.0 {
				controllers.Error(c, map[string]string{"col": "weight"}, &controllers.Msg{Str: "Please set your weight wisely"})
				return
			}
			cgdata["weight"] = id
			rsdata["weight"] = id
		}
	}

	//修改收入
	income := c.PostForm("income")
	if income != "" {
		id32, _ := strconv.Atoi(income)
		id := int64(id32)
		if id != user.Income {
			name, ok := models.IncomesMap[id]
			if !ok {
				controllers.Error(c, map[string]string{"col": "income"}, &controllers.Msg{Str: "Please set reasonably"})
				return
			}
			cgdata["income"] = id
			rsdata["income"] = name
		}
	}

	//修改工作
	job := c.PostForm("job")
	if job != "" && job != user.Job {
		if len(job) > 100 {
			controllers.Error(c, map[string]string{"col": "job"}, &controllers.Msg{Str: "Please fill in less than {{$1}} words", Args: []interface{}{100}})
			return
		}
		cgdata["job"] = job
		rsdata["job"] = job
	}

	//修改签名
	signature := c.PostForm("signature")
	if signature != "" && signature != user.Signature {
		if len(signature) > 250 {
			controllers.Error(c, map[string]string{"col": "signature"}, &controllers.Msg{Str: "Please fill in less than {{$1}} words", Args: []interface{}{250}})
			return
		}
		cgdata["signature"] = signature
		rsdata["signature"] = signature
	}

	//修改昵称
	nickname := c.PostForm("nickname")
	if nickname != "" && nickname != user.Nickname {
		if len(nickname) > 100 {
			controllers.Error(c, map[string]string{"col": "nickname"}, &controllers.Msg{Str: "Please fill in less than {{$1}} words", Args: []interface{}{100}})
			return
		}
		cgdata["nickname"] = nickname
		rsdata["nickname"] = nickname
	}

	// //修改性别
	sex := c.PostForm("sex")
	if sex != "" {
		id, _ := strconv.Atoi(sex)
		if id != user.Sex {
			if id > 2 || id < 0 {
				controllers.Error(c, map[string]string{"col": "sex"}, &controllers.Msg{Str: "Please set reasonably"})
				return
			}
			sex2name := "Confidential"
			if id == 1 {
				sex2name = "Male"
			} else if id == 2 {
				sex2name = "Female"
			}
			cgdata["sex"] = id
			rsdata["sex"] = i18n.T(lang, sex2name)
		}
	}

	// //修改风格
	temperament := c.PostForm("temperament")
	if temperament != "" {
		tst, err := user.TemSet(lang, temperament)
		if err == nil {
			if tst != nil && len(tst) > 0 {
				var toDb []string
				var toRes []string
				for _, v := range tst {
					toDb = append(toDb, fmt.Sprintf("%d", v.Id))
					toRes = append(toRes, v.Name)
				}
				cgdata["temperament"] = strings.Join(toDb, ",")
				rsdata["temperament"] = strings.Join(toRes, ",")
			} else {
				controllers.Error(c, map[string]string{"col": "temperament"}, &controllers.Msg{Str: "Please set reasonably"})
				return
			}
		}
	}

	//以下两个涉及文件上传的,必须在最后进行判断,否则将有多余的冗余
	//头像
	avatar := c.PostForm("avatar")
	if avatar != "" {
		filename, err := images.SaveFileBase64(models.AVATARPATH, fmt.Sprintf("%d%s", time.Now().Unix(), user.Invite), avatar, nil, nil)
		if err != nil {
			logs.Logger.Error(err, " - when SetAvatar model upload from form file in batch!")
			controllers.Error(c, map[string]string{"col": "avatar"}, &controllers.Msg{Str: "Image upload failed"})
			return
		}
		cgdata["avatar"] = filename
		rsdata["avatar"] = images.FullPath(filename)
	} else if avatarFile, err := c.FormFile("avatar"); err == nil {
		filename, err := images.SaveFileFromUpload(models.AVATARPATH, fmt.Sprintf("%d%s", time.Now().Unix(), user.Invite), avatarFile, nil, nil)
		if err != nil {
			logs.Logger.Error(err, " - when SetAvatar model upload from form file in batch!")
			controllers.Error(c, map[string]string{"col": "avatar"}, &controllers.Msg{Str: "Image upload failed"})
			return
		}
		cgdata["avatar"] = filename
		rsdata["avatar"] = images.FullPath(filename)
	}

	//背景图
	background := c.PostForm("background")
	if background != "" && strings.Contains(background, "base64") == true {
		filename, err := images.SaveFileBase64(models.BACKGROUNDPATH, fmt.Sprintf("%d%s", time.Now().Unix(), user.Invite), background, nil, nil)
		if err != nil {
			logs.Logger.Error(err, " - when SetBackground model upload from form file in batch!")
			controllers.Error(c, map[string]string{"col": "background"}, &controllers.Msg{Str: "Image upload failed"})
			return
		}
		cgdata["background"] = filename
		rsdata["background"] = images.FullPath(filename)
	} else if backgroundFile, err := c.FormFile("background"); err == nil {
		filename, err := images.SaveFileFromUpload(models.BACKGROUNDPATH, fmt.Sprintf("%d%s", time.Now().Unix(), user.Invite), backgroundFile, nil, nil)
		if err != nil {
			logs.Logger.Error(err, " - when SetBackground model upload from form file in batch!")
			controllers.Error(c, map[string]string{"col": "background"}, &controllers.Msg{Str: "Image upload failed"})
			return
		}
		cgdata["background"] = filename
		rsdata["background"] = images.FullPath(filename)
	}

	if cgdata != nil && len(cgdata) > 0 {
		rs := models.DB.Table("users").Where("id = ?", user.Id).Updates(cgdata)
		if rs.Error != nil {
			controllers.Error(c, nil, &controllers.Msg{Str: "Modification failure"})
			return
		}
		_, ok := cgdata["background"]
		if ok {
			os.Remove(user.Background)
		}
		_, ok = cgdata["avatar"]
		if ok {
			os.Remove(user.Avatar)
		}
		controllers.Success(c, rsdata, &controllers.Msg{Str: "Success"})
	} else {
		cs, ok := cgdata["background"]
		if ok {
			os.Remove(cs.(string))
		}
		cs, ok = cgdata["avatar"]
		if ok {
			os.Remove(cs.(string))
		}
		controllers.Success(c, nil, &controllers.Msg{Str: "No changes"})
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
		err := coder.Verify(email, code)
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

//获取签名
func Signa(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)

	controllers.Success(c, map[string]string{"signature": user.ImSignature()}, &controllers.Msg{Str: "Success"})
}

//获取用户的邀请人信息,也就是上级
func Invitee(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)

	if user.Pid > 0 {
		parent := models.GetUser(user.Pid, "", "", "")
		if parent != nil && parent.Id > 0 {
			controllers.SuccessStr(c, parent.Abstract(), "Success")
			return
		}
	}
	controllers.ErrorNoData(c, "No results found")
}

//获取被邀请人列表,也就是下级
func Invitees(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)

	lango, _ := c.Get("_lang")
	lang := lango.(string)

	timezoneo, _ := c.Get("_timezone")
	timezone := timezoneo.(string)

	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	q := c.Query("q")
	ord := c.Query("ord")

	dt, count := user.GetUserInvisList(page, limit, q, ord)
	var ddtt []map[string]interface{}
	for _, v := range dt {
		rs := v.Abstract()
		rs["addtimefmt"] = v.FmtAddTime(lang, timezone)
		ddtt = append(ddtt, rs)
	}
	controllers.SuccessStr(c, map[string]interface{}{"list": ddtt, "count": count}, "Success")
}

//账号永久注销,注销账号必须验证邮箱,否则不安全
func Cancellation(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)

	code := c.PostForm("code")
	if code == "" {
		controllers.Error(c, nil, &controllers.Msg{Str: "Please input your Captcha"})
		return
	}
	err := coder.Verify(user.Mail, code)
	if err != nil {
		controllers.Error(c, nil, &controllers.Msg{Str: err.Error()})
		return
	}

	//修改删除关系网,发送告别邮件
	bye.Send()
}
