package models

//uid 从39350开始,邀请码后4位有效,前方2位补0
import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
	"ucenter/app/config"
	"ucenter/app/funcs"
	"ucenter/app/im"
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
	Singleid      int64
	Edinfo        Editers
}

//账号修改器
type Editers int64

//创建新用户
//新用户创建必须使用 account 或 email 其中之一注册
//使用 account 必须设置密码, 使用 email 必须使用验证码
func MakeUser(account, email, phone, pwd, code, invite, nickname, platform, ip, timezone string) (user *UserModel, err error) {
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
		} else if pwd == "" {
			err = errors.New("Please set a password")
			return
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
		insertData["pid"], insertData["chain"] = GetUserInviInfo(invite)
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
	insertData["timezone"] = timezone
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
		DB.Table("users").Where("id = ?", user.Id).Update("invite", string(invicode.Encode(uint64(user.Id))))
	}

	return
}

//获取邀请用户信息
func GetUserInviInfo(invicodestr string) (pid int64, chian string) {
	invi := invicode.Decode(invicodestr)
	inviId := int64(invi)
	user := GetUser(inviId, "", "", "")
	if user != nil && user.Id > 0 {
		uschain := strings.Split(user.Chain, ",")
		uschain = append(uschain, fmt.Sprintf("%d", user.Id))
		return user.Id, strings.Join(uschain, ",")
	} else {
		return 0, ""
	}
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
	} else if limit > 100 {
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
	user := UnToken(token)

	if user.Platform < 0 { //如果通过网页注册的没用platform,则在获取用户信息时自动更新platform
		platform := c.GetHeader("platform")
		if platform != "" {

		}
	}

	return user
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
		if k == "birth" {
			if v.Int() > 0 {
				var fmt string
				fmts, ok := config.Config.Timefmts[lang]
				if ok {
					fmt = fmts.Datefmt
				} else {
					fmt = config.Config.Datefmt
				}
				data[k] = carbon.SetTimezone(timezone).CreateFromTimestamp(v.Int()).Carbon2Time().Format(fmt)
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
				data[k] = images.FullPath(s)
			} else {
				data[k] = ""
			}
		} else if k == "background" {
			s := v.String()
			if s != "" {
				data[k] = images.FullPath(s)
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
		log.Println(err, " - 初始化Im出错,请检查IM")
		return ""
	}
	str, err := imo.GenUserSig(fmt.Sprintf("%d", this.Id), 86400)
	if err != nil {
		log.Println(err, " - 获取签名错误,请检查IM")
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

//获取邀请人信息
func InviUser(code string) map[string]interface{} {
	uid := int64(invicode.Decode(code))
	if uid < 1 {
		return nil
	}
	user := GetUser(uid, "", "", "")
	if user.Id < 1 {
		return nil
	}
	dt := make(map[string]interface{})
	dt["nickname"] = user.Nickname
	dt["avatar"] = user.Avatar
	dt["id"] = user.Id
	dt["invite"] = user.Invite
	return dt
}

// //设置邀请人
// func (this *UserModel) SetInviUser(user *UserModel) {
// 	if this.Invite == "" {

// 	} else {
// 		return errors.New("No modification of invitee is allowed")
// 	}
// }
