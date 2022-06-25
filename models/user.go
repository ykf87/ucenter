package models

//uid 从39350开始,邀请码后4位有效,前方2位补0
import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
	"ucenter/app/config"
	"ucenter/app/i18n"
	"ucenter/app/mails/sender/coder"
	"ucenter/app/safety/aess"
	"ucenter/app/safety/base34"
	"ucenter/app/safety/passwordhash"
	"ucenter/app/uploadfile/images"

	"github.com/gin-gonic/gin"
	carbon "github.com/golang-module/carbon/v2"
	"github.com/oschwald/geoip2-golang"
	"github.com/tidwall/gjson"
	// "ucenter/app/safety/rsautil"
)

const (
	AVATARPATH     = "static/user/avatars/"
	BACKGROUNDPATH = "static/user/background/"
)

type UserModel struct {
	Id            int64   `json:"id"`
	Pid           int64   `json:"pid"`
	Invite        string  `json:"invite"`
	Chain         string  `json:"chain"`
	Account       string  `json:"account"`
	Mail          string  `json:"mail"`
	Phone         string  `json:"phone"`
	Mailvery      int     `json:"mailvery"`
	Phonevery     int     `json:"phonevery"`
	Pwd           string  `json:"pwd"`
	Nickname      string  `json:"nickname"`
	Avatar        string  `json:"avatar"`
	Background    string  `json:"background"`
	Addtime       int64   `json:"addtime"`
	Status        int     `json:"status"`
	Sex           int     `json:"sex"`
	Height        int     `json:"height"`
	Weight        float64 `json:"weight"`
	Birth         int64   `json:"birth"`
	Age           int     `json:"age"`
	Job           string  `json:"job"`
	Income        int64   `json:"income"`
	Emotion       int64   `json:"emotion"`
	Constellation int64   `json:"constellation"`
	Edu           int64   `json:"edu"`
	Temperament   string  `json:"temperament"`
	Ip            int64   `json:"ip"`
	Country       int64   `json:"country"`
	Province      int64   `json:"province"`
	City          int64   `json:"city"`
	Lang          string  `json:"lang"`
	Timezone      string  `json:"timezone"`
	Singleid      int64
	Edinfo        Editers
}

//账号修改器
type Editers int64

//创建新用户
//新用户创建必须使用 account 或 email 其中之一注册
//使用 account 必须设置密码, 使用 email 必须使用验证码
func MakeUser(account, email, phone, pwd, code, invite, nickname, platform, ip string) (user *UserModel, err error) {
	hadUser := new(UserModel)
	insertData := make(map[string]interface{})
	insertData["nickname"] = nickname
	if account != "" {
		DB.Table("users").Where("account = ?", account).First(hadUser)
		insertData["account"] = account
		if pwd == "" {
			err = errors.New("Please set a password")
			return
		}
	} else if email != "" {
		DB.Table("users").Where("mail = ?", email).First(hadUser)
		insertData["mail"] = email
		if code != "" {
			err = coder.Verify(email, code)
			if err != nil {
				return
			}
			insertData["mailvery"] = 1
		}

		if nickname == "" {
			tmp := strings.Split(email, "@")
			insertData["nickname"] = tmp[0]
		}
	} else {
		err = errors.New("Registration failed")
		return
	}
	// else if phone != "" {
	// 	DB.Table("users").Where("phone = ?", phone).First(hadUser)
	// 	insertData["phone"] = phone
	// }
	if hadUser.Id > 0 {
		err = errors.New("User already exists, please login")
		return
	}
	if pwd != "" {
		insertData["pwd"], err = passwordhash.PasswordHash(pwd)
		if err != nil {
			return
		}
	}
	if invite != "" {
		inviteUser := new(UserModel)
		DB.Table("users").Where("invite = ?", invite).First(inviteUser)
		if inviteUser.Id > 0 {
			insertData["pid"] = inviteUser.Id
			var addChain string
			if inviteUser.Chain != "" {
				addChain = inviteUser.Chain + ","
			}
			insertData["chain"] = addChain + fmt.Sprintf("%d", inviteUser.Id)
		}
	}

	//根据ip获取国家和城市
	geodb, errs := geoip2.Open("./GeoLite2-City.mmdb")
	if errs == nil {
		defer geodb.Close()
		ip := net.ParseIP(ip)
		record, errs := geodb.City(ip)
		if errs == nil && record.Country.IsoCode != "" {
			countryObj, errs := GetCountryByIso(record.Country.IsoCode)
			if errs == nil {
				insertData["country"] = countryObj.Id
				if cv, ok := record.City.Names["en"]; ok {
					province, errs := GetProvinceByNameAndCountry(countryObj.Id, cv, "en")
					if errs == nil {
						insertData["province"] = province.Id
					}
					city, err := GetCityByNameAndCountryId(cv, countryObj.Id, "en")
					if err == nil {
						insertData["city"] = city.Id
					}
				}
			}
		}
	}

	insertData["status"] = 1
	insertData["ip"] = InetAtoN(ip)
	insertData["addtime"] = time.Now().Unix()
	rs := DB.Table("users").Create(insertData)
	if rs.Error != nil {
		err = rs.Error
	}
	user = new(UserModel)
	user.Edinfo = Editers(user.Id)
	if DB.Table("users").Where(insertData).First(user).Error == nil && user.Id > 0 {
		if user.Pid > 0 {
			go AddUserInvitee(user.Pid, user.Id)
		}
		DB.Table("users").Where("id = ?", user.Id).Update("invite", string(base34.Base34(uint64(user.Id))))
	}

	return
}

//查找单个用户
func GetUser(id int64, account, email, phone string) *UserModel {
	user := new(UserModel)
	if id > 0 {
		DB.Table("users").Where("id = ?", id).First(user)
	} else if account != "" {
		DB.Table("users").Where("account = ?", account).First(user)
	} else if email != "" {
		DB.Table("users").Where("mail = ?", email).First(user) // and mailvery = 1
	} else if phone != "" {
		DB.Table("users").Where("phone = ?", phone).First(user) // and phonevery = 1
	}
	if user.Id < 1 {
		user = nil
	} else {
		user.Edinfo = Editers(user.Id)
	}

	return user
}

//查找用户列表
func GetUserList(page, limit int, q, rd string, noids []int64) []*UserModel {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = config.Config.Limit
	}
	dbob := DB.Table("users")
	if noids != nil && len(noids) > 0 {
		dbob = dbob.Where("id not in ?", noids)
	}
	if limit > 0 {
		dbob = dbob.Limit(limit).Offset((page - 1) * limit)
	}
	if q != "" {
		dbob = dbob.Where("nickname like ?", "%"+q+"%")
	}
	if rd != "" {
		dbob = dbob.Order("RAND()")
	}
	var useslist []*UserModel
	rs := dbob.Find(&useslist)
	if rs.Error == nil {
		return useslist
	}
	return nil
}

//通过请求获取用户信息
func GetUserFromRequest(c *gin.Context) *UserModel {
	var token string
	token = c.GetHeader("token")
	if token == "" {
		token = c.GetString("token")
	}
	if token == "" {
		return nil
	}
	return UnToken(token)
}

//生成用户token
func (this *UserModel) Token() string {
	if this.Id == 0 {
		log.Println("UserModel Token - 用户实例id为0")
		return ""
	}
	sid := this.Singleid + 1
	DB.Table("users").Where("id = ?", this.Id).Update("singleid", sid)
	this.Singleid = sid
	str := fmt.Sprintf(`{"time":%d,"id":%d,"sid":%d}`, time.Now().Unix(), this.Id, this.Singleid)
	// token, err := rsautil.RsaEncrypt(str)
	// if err != nil {
	// 	log.Println("UserModel Token - ", err)
	// 	return ""
	// }
	token := aess.EcbEncrypt(str, nil)
	return base64.StdEncoding.EncodeToString([]byte(token))
	// return token
}

//通过 token 生成 user model
func UnToken(token string) *UserModel {
	// idstr, err := rsautil.RsaDecrypt(token)
	tokens, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		log.Println("untoken base64 decode err: ", err)
		return nil
	}
	idstr := aess.EcbDecrypt(string(tokens), nil)
	if idstr == "" {
		return nil
	}
	// if err != nil {
	// 	return nil
	// }
	ts := gjson.Get(idstr, "time").Int()
	id := gjson.Get(idstr, "id").Int()
	sid := gjson.Get(idstr, "sid").Int()
	if id < 1 || ts < 1 {
		return nil
	}
	if time.Now().Unix()-ts >= 86400*365 {
		return nil
	}
	user := GetUser(int64(id), "", "", "")
	if sid != user.Singleid {
		return nil
	}
	return user
}

//返回用户信息
func (this *UserModel) Info(lang, timezone string) map[string]interface{} {
	if this.Id < 1 {
		return nil
	}
	data := make(map[string]interface{})
	b, _ := json.Marshal(this)
	for k, v := range gjson.ParseBytes(b).Map() {
		if k == "pwd" || k == "status" || k == "Singleid" || k == "chain" || k == "Edinfo" || k == "ip" || k == "pid" {
			continue
		}
		if k == "birth" && v.Int() > 0 {
			var fmt string
			fmts, ok := config.Config.Timefmts[lang]
			if ok {
				fmt = fmts.Datefmt
			} else {
				fmt = config.Config.Datefmt
			}
			data[k] = carbon.CreateFromTimestamp(v.Int()).SetTimezone(timezone).Carbon2Time().Format(fmt)
		} else if k == "country" && v.Int() > 0 {
			data[k] = CountryMap.Get(lang, v.Int())
		} else if k == "province" && v.Int() > 0 {
			r, e := GetProvinceById(v.Int(), lang)
			if e == nil {
				data[k] = r.Name
			} else {
				data[k] = ""
			}
		} else if k == "city" && v.Int() > 0 {
			r, e := GetCityById(v.Int(), lang)
			if e == nil {
				data[k] = r.Name
			} else {
				data[k] = ""
			}
		} else if k == "temperament" {
			vals := strings.Split(v.String(), ",")
			var ssds []string
			for _, v := range vals {
				id, _ := strconv.Atoi(v)
				name := TemperamentMap.Get(lang, int64(id))
				if name != "" {
					ssds = append(ssds, name)
				}
			}
			data[k] = ssds //strings.Join(ssds, ",")
		} else if k == "constellation" && v.Int() > 0 {
			data[k] = ConstellationMap.Get(lang, v.Int())
		} else if k == "edu" {
			data[k] = EducationMap.Get(lang, v.Int())
		} else if k == "emotion" {
			data[k] = EmotionMap.Get(lang, v.Int())
		} else if k == "income" {
			tmp, ok := IncomesMap[v.Int()]
			if ok == true {
				data[k] = tmp
			} else {
				data[k] = ""
			}
		} else if k == "sex" {
			sex2name := "Confidential"
			id := v.Int()
			if id == 1 {
				sex2name = "Male"
			} else if id == 2 {
				sex2name = "Female"
			}
			data[k] = i18n.T(lang, sex2name)
		} else {
			data[k] = v.String()
		}
	}
	return data
}

//账号名称修改
func (this Editers) SetAccount(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	newAccount := args[0].(string)
	if newAccount == "" {
		err = errors.New("Please set the account name")
		return
	}
	if user.Account == newAccount {
		err = errors.New("No changes")
		return
	}
	u := new(UserModel)
	DB.Table("users").Where("account = ?", newAccount).First(u)
	if u.Id > 0 {
		err = errors.New("Account name already exists, please use another one")
		return
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("account", newAccount)
	if rs.Error != nil {
		err = errors.New("Modification failure")
		return
	}
	dt = map[string]interface{}{
		"account": newAccount,
	}
	// user.Account = newAccount
	return
}

//邮箱修改
func (this Editers) SetEmail(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	newAccount := args[0].(string)
	dt = make(map[string]interface{})
	if newAccount == "" {
		err = errors.New("Please set the Email")
		return
	}
	if user.Mail == newAccount {
		err = errors.New("No changes")
		return
	}
	c := args[1].(*gin.Context)
	code := c.PostForm("code")
	if code == "" {
		err = errors.New("Please input your Captcha")
		return
	}
	err = coder.Verify(newAccount, code)
	if err != nil {
		return
	}

	u := new(UserModel)
	DB.Table("users").Where("mail = ?", newAccount).First(u)
	if u.Id > 0 {
		err = errors.New("Email already exists, please use another one")
		return
	}

	ssid := user.Singleid + 1
	ud := map[string]interface{}{
		"mail":     newAccount,
		"singleid": ssid,
		"mailvery": 1,
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Updates(ud)
	if rs.Error != nil {
		err = errors.New("Modification failure")
		return
	}
	user.Mail = newAccount
	user.Singleid = ssid
	token := user.Token()
	if token == "" {
		err = errors.New("Voucher generation failed, please try again later")
	} else {
		dt["token"] = token
	}
	return
}

//手机号修改,未接入手机证码
func (this Editers) SetPhone(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	newAccount := args[0].(string)
	if newAccount == "" {
		err = errors.New("Please set the Phone")
		return
	}
	if user.Mail == newAccount {
		err = errors.New("No changes")
		return
	}
	u := new(UserModel)
	DB.Table("users").Where("phone = ?", newAccount).First(u)
	if u.Id > 0 {
		err = errors.New("Phone already exists, please use another one")
		return
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("phone", newAccount)
	if rs.Error != nil {
		err = errors.New("Modification failure")
		return
	}
	dt = map[string]interface{}{
		"phone": newAccount,
	}
	// user.Phone = newAccount
	return
}

//修改密码
func (this Editers) SetPassword(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	changeto := args[0].(string)
	if changeto == "" {
		err = errors.New("Please set the content to be modified")
		return
	}
	npwd, _ := passwordhash.PasswordHash(changeto)
	dtn := &map[string]interface{}{
		"pwd":      npwd,
		"singleid": user.Singleid + 1,
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Updates(dtn)
	if rs.Error != nil {
		err = errors.New("Modification failure")
		return
	}
	dt = map[string]interface{}{"token": user.Token()}
	return
}

//修改昵称
func (this Editers) SetNickname(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	changeto := args[0].(string)
	if changeto == "" {
		err = errors.New("Please set the content to be modified")
		return
	}
	if len(changeto) > 200 {
		err = errors.New("Please set reasonably")
		return
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("nickname", changeto)
	if rs.Error != nil {
		err = errors.New("Modification failure")
		return
	}
	// user.Nickname = changeto
	dt = map[string]interface{}{
		"nickname": changeto,
	}
	return
}

//修改头像
func (this Editers) SetAvatar(user *UserModel, args ...interface{}) (error, map[string]interface{}) {
	changeto := args[0].(string)
	avatarPath := AVATARPATH
	var filename string
	var err error
	oldFilename := user.Avatar
	if changeto == "" {
		c := args[1].(*gin.Context)
		f, err := c.FormFile("file")
		if err != nil {
			return errors.New("Please set the content to be modified"), nil
		}
		filename, err = images.SaveFileFromUpload(avatarPath, user.Invite, f)
		if err != nil {
			log.Println(err, " - when SetAvatar model upload from form file")
			return errors.New("System error, please try again later"), nil
		}
	} else {
		filename, err = images.SaveFileBase64(avatarPath, user.Invite, changeto)
		if err != nil {
			log.Println(err, " - when SetAvatar model upload from form file")
			return errors.New("System error, please try again later"), nil
		}
	}

	if filename != "" {
		rs := DB.Table("users").Where("id = ?", user.Id).Update("avatar", filename)
		if rs.Error != nil {
			return errors.New("Modification failure"), nil
		}
		if strings.Contains(oldFilename, filename) != true {
			os.Remove(oldFilename)
		}
		// user.Avatar = filename
		return nil, map[string]interface{}{"avatar": config.Config.Domain + filename}
	}
	return errors.New("System error, please try again later"), nil
}

//修改背景图片
func (this Editers) SetBackground(user *UserModel, args ...interface{}) (error, map[string]interface{}) {
	changeto := args[0].(string)
	avatarPath := BACKGROUNDPATH
	var filename string
	var err error
	oldFilename := user.Background
	if changeto == "" {
		c := args[1].(*gin.Context)
		f, err := c.FormFile("file")
		if err != nil {
			return errors.New("Please set the content to be modified"), nil
		}
		filename, err = images.SaveFileFromUpload(avatarPath, user.Invite, f)
		if err != nil {
			log.Println(err, " - when SetBackground model upload from form file")
			return errors.New("System error, please try again later"), nil
		}
	} else {
		filename, err = images.SaveFileBase64(avatarPath, user.Invite, changeto)
		if err != nil {
			log.Println(err, " - when SetAvatar model upload from form file")
			return errors.New("System error, please try again later"), nil
		}
	}

	if filename != "" {
		rs := DB.Table("users").Where("id = ?", user.Id).Update("background", filename)
		if rs.Error != nil {
			return errors.New("Modification failure"), nil
		}
		if strings.Contains(oldFilename, filename) != true {
			os.Remove(oldFilename)
		}
		// user.Background = filename
		return nil, map[string]interface{}{"background": filename}
	}
	return errors.New("System error, please try again later"), nil
}

//修改性别
func (this Editers) SetSex(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	changeto := args[0].(string)
	if changeto == "" {
		err = errors.New("Please set the content to be modified")
		return
	}
	if changeto != "0" && changeto != "1" && changeto != "2" {
		changeto = "0"
	}
	cgt, _ := strconv.Atoi(changeto)
	rs := DB.Table("users").Where("id = ?", user.Id).Update("sex", cgt)
	if rs.Error != nil {
		err = errors.New("Modification failure")
		return
	}
	dt = map[string]interface{}{
		"sex": cgt,
	}
	// user.Sex = cgt
	return
}

//修改身高
func (this Editers) SetHeight(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	changeto := args[0].(string)
	if changeto == "" {
		err = errors.New("Please set the content to be modified")
		return
	}
	cgt, _ := strconv.Atoi(changeto)
	if cgt < 1 || cgt >= 300 {
		err = errors.New("Please set the height reasonably")
		return
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("height", cgt)
	if rs.Error != nil {
		err = errors.New("Modification failure")
		return
	}
	dt = map[string]interface{}{
		"height": cgt,
	}
	// user.Height = cgt
	return
}

//修改体重
func (this Editers) SetWeight(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	changeto := args[0].(string)
	if changeto == "" {
		err = errors.New("Please set the content to be modified")
		return
	}
	cgt, _ := strconv.ParseFloat(changeto, 64)
	if cgt < 1.0 || cgt >= 500000.0 {
		err = errors.New("Please set your weight wisely")
		return
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("weight", cgt)
	if rs.Error != nil {
		err = errors.New("Modification failure")
		return
	}
	dt = map[string]interface{}{
		"weight": cgt,
	}
	// user.Weight = cgt
	return
}

//修改年龄
func (this Editers) SetAge(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	changeto := args[0].(string)
	if changeto == "" {
		err = errors.New("Please set the content to be modified")
		return
	}
	cgt, _ := strconv.Atoi(changeto)
	if cgt < 1 || cgt > 200 {
		err = errors.New("Wrong age")
		return
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("age", cgt)
	if rs.Error != nil {
		err = errors.New("Modification failure")
		return
	}
	dt = map[string]interface{}{
		"age": cgt,
	}
	// user.Age = cgt
	return
}

//修改生日
func (this Editers) SetBirth(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	changeto := args[0].(string)
	if changeto == "" {
		err = errors.New("Please set the content to be modified")
		return
	}
	cgt := carbon.Parse(changeto).Timestamp()
	if cgt == 0 {
		err = errors.New("Wrong date format")
		return
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("birth", cgt)
	if rs.Error != nil {
		err = errors.New("Modification failure")
		return
	}
	dt = map[string]interface{}{
		"birth": changeto,
	}
	// user.Birth = cgt
	return
}

//修改工作
func (this Editers) SetJob(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	changeto := args[0].(string)
	if changeto == "" {
		err = errors.New("Please set the content to be modified")
		return
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("job", changeto)
	if rs.Error != nil {
		err = errors.New("Modification failure")
		return
	}
	// user.Job = changeto
	dt = map[string]interface{}{
		"job": changeto,
	}
	return
}

//修改收入
func (this Editers) SetIncome(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	changeto := args[0].(string)
	if changeto == "" {
		err = errors.New("Please set the content to be modified")
		return
	}
	cgt, _ := strconv.Atoi(changeto)
	if cgt < 1 {
		err = errors.New("Please set reasonably")
		return
	}
	ccgt := int64(cgt)
	ins, ok := IncomesMap[ccgt]
	if !ok {
		err = errors.New("Please set reasonably")
		return
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("income", ccgt)
	if rs.Error != nil {
		err = errors.New("Modification failure")
		return
	}
	dt = map[string]interface{}{
		"income": ins,
	}
	// user.Income = ccgt
	return
}

//修改情感状态
func (this Editers) SetEmotion(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	changeto := args[0].(string)
	if changeto == "" {
		err = errors.New("Please set the content to be modified")
		return
	}

	c := args[1].(*gin.Context)
	langob, _ := c.Get("_lang")
	lang := langob.(string)

	id32, _ := strconv.Atoi(changeto)
	id := int64(id32)
	name := EmotionMap.Get(lang, id)
	if name == "" {
		err = errors.New("Please set reasonably")
		return
	}

	rs := DB.Table("users").Where("id = ?", user.Id).Update("emotion", id)
	if rs.Error != nil {
		err = errors.New("Modification failure")
		return
	}
	dt = map[string]interface{}{
		"emotion": name,
	}
	// user.Emotion = id
	return
}

//修改星座
func (this Editers) SetConstellation(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	changeto := args[0].(string)
	if changeto == "" {
		err = errors.New("Please set the content to be modified")
		return
	}

	c := args[1].(*gin.Context)
	langob, _ := c.Get("_lang")
	lang := langob.(string)

	id32, _ := strconv.Atoi(changeto)
	id := int64(id32)
	name := ConstellationMap.Get(lang, id)
	if name == "" {
		err = errors.New("Please set reasonably")
		return
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("constellation", id)
	if rs.Error != nil {
		err = errors.New("Modification failure")
		return
	}
	// user.Constellation = id
	dt = map[string]interface{}{
		"constellation": name,
	}
	return
}

//修改性格
func (this Editers) SetTemperament(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	changeto := args[0].(string)
	if changeto == "" {
		err = errors.New("Please set the content to be modified")
		return
	}

	c := args[1].(*gin.Context)
	langob, _ := c.Get("_lang")
	lang := langob.(string)

	var vals []string
	var ids []string
	arrs := strings.Split(changeto, ",")
	for _, v := range arrs {
		cgt, _ := strconv.Atoi(v)
		geted := TemperamentMap.Get(lang, int64(cgt))
		if geted != "" {
			vals = append(vals, geted)
			ids = append(ids, v)
		}
	}

	if len(ids) > 0 {
		changeto = strings.Join(ids, ",")
		rs := DB.Table("users").Where("id = ?", user.Id).Update("temperament", changeto)
		if rs.Error != nil {
			err = errors.New("Modification failure")
			return
		}
		// user.Temperament = changeto
	} else {
		err = errors.New("Please set reasonably")
		return
	}

	dt = map[string]interface{}{
		"constellation": vals,
	}
	return
}

//修改学历
func (this Editers) SetEdu(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	changeto := args[0].(string)
	if changeto == "" {
		err = errors.New("Please set the content to be modified")
		return
	}
	id32, _ := strconv.Atoi(changeto)
	id := int64(id32)
	if id < 1 {
		err = errors.New("Please set the content to be modified")
		return
	}
	if user.Edu == id {
		err = errors.New("No changes")
		return
	}

	c := args[1].(*gin.Context)
	langob, _ := c.Get("_lang")
	lang := langob.(string)

	eduname := EducationMap.Get(lang, id)
	if eduname == "" {
		err = errors.New("Please set the content to be modified")
		return
	}

	rs := DB.Table("users").Where("id = ?", user.Id).Update("edu", id)
	if rs.Error != nil {
		err = errors.New("Modification failure")
		return
	}
	// user.Edu = id
	dt = map[string]interface{}{"edu": eduname}
	return
}

//修改国家
func (this Editers) SetCountry(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	changeto := args[0].(string)
	if changeto == "" {
		err = errors.New("Please set the content to be modified")
		return
	}
	id, _ := strconv.Atoi(changeto)
	if id < 1 {
		err = errors.New("Please set the content to be modified")
		return
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("country", id)
	if rs.Error != nil {
		err = errors.New("Modification failure")
		return
	}
	// user.Country = int64(id)
	c := args[1].(*gin.Context)
	langob, _ := c.Get("_lang")
	lang := langob.(string)
	dt = map[string]interface{}{
		"country": CountryMap.Get(lang, int64(id)),
	}
	return
}

//修改城市
func (this Editers) SetCity(user *UserModel, args ...interface{}) (err error, dt map[string]interface{}) {
	changeto := args[0].(string)
	if changeto == "" {
		err = errors.New("Please set the content to be modified")
		return
	}
	if user.Country < 1 {
		err = errors.New("Please select your country first")
		return
	}
	cityid32, _ := strconv.Atoi(changeto)
	cityid := int64(cityid32)
	if user.City == cityid {
		err = errors.New("No changes")
		return
	}

	c := args[1].(*gin.Context)
	langob, _ := c.Get("_lang")
	citylist, errs := GetCityByCountryId(langob.(string), user.Country)
	if err != nil {
		err = errs
		return
	}
	if len(citylist) < 1 {
		err = errors.New("Please select your country first")
		return
	}
	var findcity *CityModel
	for _, v := range citylist {
		if v.Id == cityid {
			findcity = v
			break
		}
	}
	if findcity == nil {
		err = errors.New("The city of your choice was not found")
		return
	}

	if cityid < 1 {
		err = errors.New("Please set the content to be modified")
		return
	}
	rs := DB.Table("users").Where("id = ?", user.Id).Update("city", cityid)
	if rs.Error != nil {
		err = errors.New("Modification failure")
		return
	}
	// user.City = int64(cityid)
	dt = map[string]interface{}{"city": findcity.Name}
	return
}

//修改密码
func (this *UserModel) ChangePwd(pwd string) error {
	if pwd == "" {
		return errors.New("Please set a password")
	}
	pwds, err := passwordhash.PasswordHash(pwd)
	if err != nil {
		return err
	}
	tosign := this.Singleid + 1
	dt := map[string]interface{}{
		"pwd":      pwds,
		"singleid": tosign,
	}
	rs := DB.Table("users").Where("id = ?", this.Id).Updates(dt)
	if rs.Error != nil {
		this.Pwd = pwds
		this.Singleid = tosign
		return nil
	} else {
		return rs.Error
	}
}

//注销账号,账号信息存储至其他表格,并删除user表内容
func (this *UserModel) MoveAndDelete() {

}
