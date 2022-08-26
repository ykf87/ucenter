package models

//uid 从39350开始,邀请码后4位有效,前方2位补0
import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
	"ucenter/app/config"
	"ucenter/app/funcs"
	"ucenter/app/im"
	"ucenter/app/logs"
	"ucenter/app/mails/sender/coder"
	"ucenter/app/safety/aess"
	"ucenter/app/safety/invicode"
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
	REALUSERPATH   = "static/user/realuser/"
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
	Signature     string  `json:"signature"`
	Visits        int64   `json:"visits"`
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
	Platform      int     `json:"platform"`
	Md5           string  `json:"md5"`
	Realuser      string  `json:"realuser"`
	Singleid      int64   `json:"singleid"`
	Recharge      float64 `json:"-"`
	Used          float64 `json:"-"`
	Balance       float64 `json:"-"`
}

//账号修改器

//创建新用户
//新用户创建必须使用 account 或 email 其中之一注册
//使用 account 必须设置密码, 使用 email 必须使用验证码
func MakeUser(account, email, phone, pwd, code, invite, nickname, platform, ip, timezone string) (user *UserModel, err error) {
	hadUser := new(UserModel)
	// insertData := make(map[string]interface{})
	insertUser := new(UserModel)
	insertUser.Nickname = nickname
	if account != "" {
		DB.Table("users").Where("account = ?", account).Where("status >= 0").First(hadUser)
		insertUser.Account = account
		if pwd == "" {
			err = errors.New("Please set a password")
			return
		}
	} else if email != "" {
		DB.Table("users").Where("mail = ?", email).Where("status >= 0").First(hadUser)
		insertUser.Mail = email
		if code != "" {
			err = coder.Verify(email, code)
			if err != nil {
				return
			}
			insertUser.Mailvery = 1
		} else if pwd == "" {
			err = errors.New("Please set a password")
			return
		}

		if nickname == "" {
			tmp := strings.Split(email, "@")
			insertUser.Nickname = tmp[0]
		}
	} else {
		err = errors.New("Registration failed")
		return
	}
	// else if phone != "" {
	// 	DB.Table("users").Where("phone = ?", phone).Where("status >= 0").First(hadUser)
	// 	insertData["phone"] = phone
	// }
	if hadUser.Id > 0 {
		err = errors.New("User already exists, please login")
		return
	}
	if pwd != "" {
		insertUser.Pwd, err = passwordhash.PasswordHash(pwd)
		if err != nil {
			return
		}
	}
	if invite != "" {
		parent := GetUserInviInfo(invite)
		if parent != nil && parent.Id > 0 {
			var chianArr []string
			if parent.Chain != "" {
				chianArr = strings.Split(parent.Chain, ",")
			}
			chianArr = append(chianArr, fmt.Sprintf("%d", parent.Id))
			insertUser.Pid, insertUser.Chain = parent.Id, strings.Join(chianArr, ",")
		}

		// DB.Table("users").Where("invite = ?", invite).First(inviteUser)
		// if inviteUser.Id > 0 {

		// 	insertData["pid"] = inviteUser.Id
		// 	var addChain string
		// 	if inviteUser.Chain != "" {
		// 		addChain = inviteUser.Chain + ","
		// 	}
		// 	insertData["chain"] = addChain + fmt.Sprintf("%d", inviteUser.Id)
		// }
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
				insertUser.Country = countryObj.Id
				if countryObj.Timezone != "" {
					insertUser.Timezone = timezone
				}
				if cv, ok := record.City.Names["en"]; ok {
					province, errs := GetProvinceByNameAndCountry(countryObj.Id, cv, "en")
					if errs == nil {
						insertUser.Province = province.Id
					}
					city, err := GetCityByNameAndCountryId(cv, countryObj.Id, "en")
					if err == nil {
						insertUser.City = city.Id
					}
				}
			}
		}
	} else {
		logs.Logger.Error(errs)
	}

	insertUser.Status = 1
	insertUser.Ip = funcs.InetAtoN(ip)
	insertUser.Addtime = time.Now().Unix()
	rs := DB.Table("users").Create(insertUser)
	if rs.Error != nil {
		err = rs.Error
	}
	if insertUser.Id > 0 {
		user = insertUser
		if user.Pid > 0 {
			go AddUserInvitee(user.Pid, user.Id)
		}
		DB.Table("users").Where("id = ?", user.Id).Update("invite", string(invicode.Encode(uint64(user.Id))))
	}

	return
}

//查找单个用户
func GetUser(id int64, account, email, phone string) *UserModel {
	tbName := "users"
	user := new(UserModel)
	if id > 0 {
		DB.Table(tbName).Where("id = ?", id).Where("status >= 0").First(user)
	} else if account != "" {
		// DB.Table(tbName).Where("account = ?", account).Where("status >= 0").First(user)
	} else if email != "" {
		DB.Table(tbName).Where("mail = ?", email).Where("status >= 0").First(user) // and mailvery = 1
	} else if phone != "" {
		// DB.Table(tbName).Where("phone = ?", phone).First(user) // and phonevery = 1
	}

	if user.Id < 1 {
		return nil
	}
	return user
}

//查找用户列表
func GetUserList(page, limit int, q, rd string, noids []int64, searcherSex int) []*UserModel {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = config.Config.Limit
	} else if limit > 100 {
		limit = config.Config.Limit
	}
	dbob := DB.Table("users").Where("sex != 0").Where("status >= 0")
	if noids != nil && len(noids) > 0 {
		dbob = dbob.Where("id not in ?", noids)
	}
	if limit > 0 {
		dbob = dbob.Limit(limit).Offset((page - 1) * limit)
	}
	if q != "" {
		dbob = dbob.Where("nickname like ?", "%"+q+"%")
	}
	if searcherSex > 0 && config.Config.Heterosexual == 1 {
		dbob = dbob.Where("sex != ?", searcherSex)
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

//查询用户邀请列表
func (this *UserModel) GetUserInvisList(page, limit int, q, ord string) ([]*UserModel, int64) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = config.Config.Limit
	} else if limit > 100 {
		limit = config.Config.Limit
	}
	dbs := DB.Select("a.*").Table("users as a").Joins("left join user_invitees as b on a.id = b.uid").Where("b.id = ?", this.Id).Where("a.status >= 0")
	if q != "" {
		dbs = dbs.Where("a.nickname like ?", "%"+q+"%")
	}

	dbs = dbs.Limit(limit).Offset((page - 1) * limit)

	switch ord {
	case "addasc":
		dbs = dbs.Order("a.addtime ASC")
	default:
		dbs = dbs.Order("a.addtime DESC")
	}

	var list []*UserModel
	var total int64
	dbs.Find(&list)
	if page == 1 && len(list) > 0 {
		DB.Table("users as a").Joins("left join user_invitees as b on a.id = b.uid").Where("b.id = ?", this.Id).Count(&total)
	}
	return list, total
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
	user := UnToken(token)

	if user != nil && user.Id > 0 && user.Platform < 0 { //如果通过网页注册的没用platform,则在获取用户信息时自动更新platform
		platform := c.GetHeader("platform")
		platid, _ := strconv.Atoi(platform)
		if platid > 0 && platid < 10 {
			DB.Table("users").Where("id = ?", user.Id).Update("platform", platid)
		}
	}

	return user
}

//生成用户token
func (this *UserModel) Token() string {
	if this.Id == 0 {
		logs.Logger.Error("UserModel Token - 用户实例id为0")
		return ""
	}
	sid := this.Singleid + 1
	DB.Table("users").Where("id = ?", this.Id).Update("singleid", sid)
	this.Singleid = sid
	str := fmt.Sprintf(`{"time":%d,"id":%d,"sid":%d}`, time.Now().Unix(), this.Id, this.Singleid)

	token := aess.EcbEncrypt(str, nil)

	go AddUserLoginRow(this)

	return base64.StdEncoding.EncodeToString([]byte(token))
}

//通过 token 生成 user model
func UnToken(token string) *UserModel {
	// idstr, err := rsautil.RsaDecrypt(token)
	tokens, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		logs.Logger.Error("untoken base64 decode err: ", err)
		return nil
	}
	idstr := aess.EcbDecrypt(string(tokens), nil)
	if idstr == "" {
		return nil
	}
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
	if user == nil || sid != user.Singleid {
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
		if k == "birth" {
			if v.Int() > 0 {
				data[k] = this.FmtAddTime(lang, timezone)
			} else {
				data[k] = ""
			}
		} else if k == "country" {
			if v.Int() > 0 {
				data[k] = CountryMap.Get(lang, v.Int())
				data["countryid"] = v.Int()
			} else {
				data[k] = ""
			}
		} else if k == "province" {
			if v.Int() > 0 {
				r, e := GetProvinceById(v.Int(), lang)
				if e == nil {
					data[k] = r.Name
				} else {
					data[k] = ""
				}
			} else {
				data[k] = ""
			}
		} else if k == "city" {
			if v.Int() > 0 {
				r, e := GetCityById(v.Int(), lang)
				if e == nil {
					data[k] = r.Name
				} else {
					data[k] = ""
				}
			} else {
				data[k] = ""
			}
		} else if k == "temperament" {
			vals := strings.Split(v.String(), ",")
			var ssds []string
			for _, v := range vals {
				id32, _ := strconv.Atoi(v)
				id := int64(id32)
				name := TemperamentMap.Get(lang, id)
				if name != "" {
					ssds = append(ssds, name)
				}
			}
			data[k] = strings.Join(ssds, ",")
		} else if k == "constellation" {
			if v.Int() > 0 {
				data[k] = ConstellationMap.Get(lang, v.Int())
			} else {
				data[k] = ""
			}
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
		} else if k == "avatar" {
			s := v.String()
			if s != "" {
				data[k] = images.FullPath(s, "")
				data["avatar_thumb"] = images.FullPath(s, "small")
			} else {
				data[k] = ""
			}
		} else if k == "background" {
			s := v.String()
			if s != "" {
				data[k] = images.FullPath(s, "")
				data["background_thumb"] = images.FullPath(s, "medium")
			} else {
				data[k] = ""
			}
		} else {
			data[k] = v.String()
		}
	}
	return data
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

//获取用户im签名
func (this *UserModel) ImSignature() string {
	imo, err := im.Get(config.Config.Useim)
	if err != nil {
		logs.Logger.Error(err, " - 初始化Im出错,请检查IM")
		return ""
	}
	str, err := imo.GenUserSig(fmt.Sprintf("%d", this.Id), 86400)
	if err != nil {
		logs.Logger.Error(err, " - 获取签名错误,请检查IM")
		return ""
	}
	return str
}

//修改风格
func (this *UserModel) TemSet(lang, str string) (res []*IdNameModel, err error) {
	if str == "" {
		err = errors.New("Please set the content to be modified")
		return
	}
	//先对str去重
	strs := funcs.RemoveRepByMap(strings.Split(str, ","))

	newTemp := this.temSet(lang, strs)
	var userTemp []*IdNameModel
	if this.Temperament != "" {
		userTemp = this.temSet(lang, strings.Split(this.Temperament, ","))
	}
	if userTemp != nil && len(userTemp) > 0 { //如果原来有设置,则要验证是否重复
		var newTotal, userTotal int64
		for _, v := range newTemp {
			newTotal = newTotal + v.Id
		}
		for _, v := range userTemp {
			userTotal = userTotal + v.Id
		}
		if userTotal == newTotal {
			err = errors.New("No changes")
			return
		}
		res = newTemp
	} else {
		res = newTemp
	}
	return
}
func (this *UserModel) temSet(lang string, strs []string) []*IdNameModel {
	var ress []*IdNameModel
	for _, v := range strs {
		id32, _ := strconv.Atoi(v)
		id := int64(id32)
		name := TemperamentMap.Get(lang, id)
		if name != "" {
			idm := new(IdNameModel)
			idm.Id = id
			idm.Name = name
			ress = append(ress, idm)
		}
	}
	return ress
}

//增加访问量
func (this *UserModel) AddVisits() {
	DB.Table("users").Where("id = ?", this.Id).Update("visits", this.Visits+1)
}

//注销账号,账号信息存储至其他表格,并删除user表内容
func (this *UserModel) MoveAndDelete() {

}

//获取邀请用户信息
func GetUserInviInfo(invicodestr string) *UserModel {
	invi := invicode.Decode(invicodestr)
	inviId := int64(invi)
	return GetUser(inviId, "", "", "")
}

//获取邀请人信息
func (this *UserModel) Abstract() map[string]interface{} {
	dt := make(map[string]interface{})
	if this.Id > 0 {
		dt["nickname"] = this.Nickname
		dt["main"] = this.Mail
		dt["avatar"] = images.FullPath(this.Avatar, "")
		dt["avatar_thumb"] = images.FullPath(this.Avatar, "small")
		dt["id"] = this.Id
		dt["invite"] = this.Invite
		dt["addtime"] = this.Addtime
		// dt["inviurl"] = funcs.InviUrl(this.Invite)
	}

	return dt
}

//格式化时间
func (this *UserModel) FmtAddTime(lang, timezone string) string {
	var fmt string
	fmts, ok := config.Config.Timefmts[lang]
	if ok {
		fmt = fmts.Datefmt
	} else {
		fmt = config.Config.Datefmt
	}
	return carbon.SetTimezone(timezone).CreateFromTimestamp(this.Addtime).Carbon2Time().Format(fmt)
}

//检查当前登录环境是否较上次有变动
func (this *UserModel) CheckUserUseEnvironment(c *gin.Context, isreg bool) {
	md5str := funcs.UserDeviceMd5(c)
	if md5str == "" {
		return
	}
	if md5str == this.Md5 {
		return
	}
	AddUserEnvironmentChange(c, this, isreg)
	DB.Table("users").Where("id = ?", this.Id).Update("md5", md5str)
	return
}

//登录成功后的返回
func (this *UserModel) UserAfterLogin() map[string]interface{} {
	ddt := make(map[string]interface{})

	ddt["token"] = this.Token()
	ddt["id"] = this.Id
	ddt["signature"] = this.ImSignature()
	return ddt
}

func (this *UserModel) Cancellation() {
	rs := DB.Table("users").Where("id = ?", this.Id).Updates(map[string]interface{}{"status": -1, "singleid": -1})
	fmt.Println(rs.Error)
}

//获取余额
func (this *UserModel) GetUserBalance() int64 {
	reched := Recharged(this.Id)
	if reched <= 0 {
		return reched
	}
	used := IUsed(this.Id)
	return reched - used
}

//更改用户充值总额
func (this *UserModel) ChangeRecharge(recharge float64) {
	if recharge <= 0 {
		r, _ := strconv.ParseFloat(fmt.Sprintf("%d", Recharged(this.Id)), 64)
		this.Recharge = r
	} else {
		this.Recharge += recharge
	}
	data := map[string]interface{}{
		"recharge": this.Recharge,
		"balance":  this.Recharge - this.Used,
	}
	DB.Table("users").Where("id = ?", this.Id).Updates(data)
}

//更改用户使用总额
func (this *UserModel) ChangeUsed(used float64, balance float64) {
	if used <= 0 && balance <= 0 {
		return
	}
	if used <= 0 {
		used = this.Recharge - balance
	}
	if balance <= 0 {
		balance = this.Recharge - used
	}

	data := map[string]float64{
		"used":    this.Used,
		"balance": balance,
	}
	DB.Model(&UserModel{}).Updates(data)
}
